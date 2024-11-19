package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

// Config holds the configuration parameters
type Config struct {
	PcapFile     string `yaml:"pcapFile"`
	IfaceName    string `yaml:"ifaceName"`
	LanguageFile string `yaml:"languageFile"`
}

// Messages holds all the messages and errors
type Messages struct {
	Errors   map[string]string `yaml:"errors"`
	Messages map[string]string `yaml:"messages"`
}

var messages Messages

func main() {
	if err := run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func run() error {
	configFile, listInterfaces := parseFlags()

	if listInterfaces {
		printInterfaces()
		return nil
	}

	config, err := readConfig(configFile)
	if err != nil {
		return err
	}

	// Load messages from language file
	if err := loadMessages(config.LanguageFile); err != nil {
		return err
	}

	if err := validateConfig(config); err != nil {
		return err
	}

	fileHandle, reader, err := openPcapFile(config.PcapFile)
	if err != nil {
		return err
	}
	defer fileHandle.Close()

	ifaceHandle, err := openNetworkInterface(config.IfaceName)
	if err != nil {
		return err
	}
	defer ifaceHandle.Close()

	fmt.Println(getMessage("messages.starting_replay", nil))
	if err := replayPackets(reader, ifaceHandle); err != nil {
		return err
	}
	fmt.Println(getMessage("messages.replay_completed", nil))

	return nil
}

func parseFlags() (string, bool) {
	var configFile string
	var listInterfaces bool

	flag.StringVar(&configFile, "c", "./Config/config.yaml", "Path to the configuration file")
	flag.BoolVar(&listInterfaces, "l", false, "List available network interfaces")
	flag.Parse()

	return configFile, listInterfaces
}

func readConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf(getMessage("errors.config_open_error", map[string]interface{}{
			"file":  configFile,
			"error": err,
		}))
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf(getMessage("errors.config_decode_error", map[string]interface{}{
			"error": err,
		}))
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	if config.PcapFile == "" {
		return fmt.Errorf(getMessage("errors.config_pcapfile_missing", nil))
	}
	if config.IfaceName == "" {
		return fmt.Errorf(getMessage("errors.config_ifacename_missing", nil))
	}
	if config.LanguageFile == "" {
		config.LanguageFile = "./Config/messages.yaml" // Default language file
	}
	return nil
}

func loadMessages(languageFile string) error {
	file, err := os.Open(languageFile)
	if err != nil {
		return fmt.Errorf("Error opening language file %s: %v", languageFile, err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&messages); err != nil {
		return fmt.Errorf("Error decoding language file: %v", err)
	}

	return nil
}

func getMessage(key string, data map[string]interface{}) string {
	var tmplStr string
	if val, ok := messages.Messages[key]; ok {
		tmplStr = val
	} else if val, ok := messages.Errors[key]; ok {
		tmplStr = val
	} else {
		return key // Return the key if message not found
	}

	tmpl, err := template.New("message").Parse(tmplStr)
	if err != nil {
		return tmplStr
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return tmplStr
	}

	return buf.String()
}

func openPcapFile(pcapFile string) (*os.File, *pcapgo.NgReader, error) {
	handle, err := os.Open(pcapFile)
	if err != nil {
		return nil, nil, fmt.Errorf(getMessage("errors.pcap_open_error", map[string]interface{}{
			"error": err,
		}))
	}

	reader, err := pcapgo.NewNgReader(handle, pcapgo.DefaultNgReaderOptions)
	if err != nil {
		handle.Close()
		return nil, nil, fmt.Errorf(getMessage("errors.pcap_reader_error", map[string]interface{}{
			"error": err,
		}))
	}

	return handle, reader, nil
}

func openNetworkInterface(ifaceName string) (*pcap.Handle, error) {
	ifaceHandle, err := pcap.OpenLive(ifaceName, 65535, false, pcap.BlockForever)
	if err != nil {
		return nil, fmt.Errorf(getMessage("errors.iface_open_error", map[string]interface{}{
			"iface": ifaceName,
			"error": err,
		}))
	}

	if err := ifaceHandle.SetDirection(pcap.DirectionOut); err != nil {
		ifaceHandle.Close()
		return nil, fmt.Errorf(getMessage("errors.iface_direction_error", map[string]interface{}{
			"iface": ifaceName,
			"error": err,
		}))
	}

	return ifaceHandle, nil
}

func replayPackets(reader *pcapgo.NgReader, ifaceHandle *pcap.Handle) error {
	packetSource := gopacket.NewPacketSource(reader, reader.LinkType())

	var firstPacketTime time.Time
	var lastPacketTime time.Time

	for packet := range packetSource.Packets() {
		packetTime := packet.Metadata().Timestamp

		if firstPacketTime.IsZero() {
			firstPacketTime = packetTime
			lastPacketTime = packetTime
		}

		delay := packetTime.Sub(lastPacketTime)
		if delay > 0 {
			time.Sleep(delay)
		}

		if err := ifaceHandle.WritePacketData(packet.Data()); err != nil {
			return fmt.Errorf(getMessage("errors.packet_send_error", map[string]interface{}{
				"error": err,
			}))
		}

		lastPacketTime = packetTime
	}

	return nil
}

// printInterfaces lists all available network interfaces
func printInterfaces() {
	ifaces, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf(getMessage("errors.interfaces_find_error", map[string]interface{}{
			"error": err,
		}))
	}

	if len(ifaces) == 0 {
		fmt.Println(getMessage("messages.no_interfaces_found", nil))
		return
	}

	fmt.Println(getMessage("messages.available_interfaces", nil))
	for _, iface := range ifaces {
		fmt.Println(getMessage("messages.iface_name", map[string]interface{}{
			"name": iface.Name,
		}))
		if iface.Description != "" {
			fmt.Println(getMessage("messages.iface_description", map[string]interface{}{
				"description": iface.Description,
			}))
		}
		for _, address := range iface.Addresses {
			fmt.Println(getMessage("messages.iface_ip_address", map[string]interface{}{
				"ip": address.IP,
			}))
		}
		fmt.Println()
	}
}
