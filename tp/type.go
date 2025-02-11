package tp

type Clash struct {
	MixedPort               int     `yaml:"mixed-port"`
	AllowLan                bool    `yaml:"allow-lan"`
	Mode                    string  `yaml:"mode"`
	LogLevel                string  `yaml:"log-level"`
	ExternalController      string  `yaml:"external-controller"`
	Secret                  string  `yaml:"secret"`
	Ipv6                    bool    `yaml:"ipv6"`
	UnifiedDelay            bool    `yaml:"unified-delay"`
	TcpConcurrent           bool    `yaml:"tcp-concurrent"`
	FindProcessMode         string  `yaml:"find-process-mode"`
	GlobalClientFingerprint string  `yaml:"global-client-fingerprint"`
	GeodataMode             bool    `yaml:"geodata-mode"`
	GeodataLoader           string  `yaml:"geodata-loader"`
	GeoxUrl                 GeoxUrl `yaml:"geox-url"`
	GeoAutoUpdate           bool    `yaml:"geo-auto-update"`
	GeoUpdateInterval       int     `yaml:"geo-update-interval"`
	Profile                 Profile `yaml:"profile"`
	Sniffer                 Sniffer `yaml:"sniffer"`
	Dns                     Dns     `yaml:"dns"`
	Tun                     Tun     `yaml:"tun"`
	Proxies                 Proxies `yaml:"proxies"`
}

type GeoxUrl struct {
	Geoip   string `yaml:"geoip"`
	Geosite string `yaml:"geosite"`
	Mmdb    string `yaml:"mmdb"`
	Asn     string `yaml:"asn"`
}

type Profile struct {
	StoreSelected bool `yaml:"store-selected"`
	StoreFakeIp   bool `yaml:"store-fake-ip"`
}

type Sniffer struct {
	Enable bool  `yaml:"enable"`
	Sniff  Sniff `yaml:"sniff"`
}

type Sniff struct {
	Http       Http       `yaml:"HTTP"`
	Tls        Tls        `yaml:"TLS"`
	Quic       Quic       `yaml:"QUIC"`
	SkipDomain SkipDomain `yaml:"skip-domain"`
}

type Ports []string

type Domains []string

type Http struct {
	Ports
}

type Tls struct {
	Ports
}
type Quic struct {
	Ports
}

type SkipDomain struct {
	Domains
}

type Dns struct {
	Enable                bool                  `yaml:"enable"`
	PreferH3              bool                  `yaml:"prefer-h3"`
	Listen                string                `yaml:"listen"`
	Ipv6                  bool                  `yaml:"ipv6"`
	RespectRules          bool                  `yaml:"respect-rules"`
	EnhancedMode          string                `yaml:"enhanced-mode"`
	FakeIpRange           string                `yaml:"fake-ip-range"`
	UseHosts              bool                  `yaml:"use-hosts"`
	FakeIpFilter          FakeIpFilter          `yaml:"fake-ip-filter"`
	Nameserver            Nameserver            `yaml:"nameserver"`
	ProxyServerNameserver ProxyServerNameserver `yaml:"proxy-server-nameserver"`
}

type FakeIpFilter struct {
	Domains
}

type Nameserver struct {
	Domains
}

type ProxyServerNameserver struct {
	Domains
}

type Tun struct {
	Enable              bool      `yaml:"enable"`
	Stack               string    `yaml:"stack"`
	Device              string    `yaml:"device"`
	AutoRoute           bool      `yaml:"auto-route"`
	AutoRedirect        bool      `yaml:"auto-redirect"`
	Mtu                 int       `yaml:"mtu"`
	AutoDetectInterface bool      `yaml:"auto-detect-interface"`
	DnsHijack           DnsHijack `yaml:"dns-hijack"`
}

type DnsHijack struct {
	Domains
}

type Proxies struct {
}

type Proxy struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	Server     string `yaml:"server"`
	Port       int    `yaml:"port"`
	Network    string `yaml:"network"`
	Tls        bool   `yaml:"tls"`
	Udp        bool   `yaml:"udp"`
	Servername string `yaml:"servername"`
}

type Reality struct {
	Uuid              string      `yaml:"uuid"`
	Flow              string      `yaml:"flow"`
	RealityOpts       RealityOpts `yaml:"reality-opts"`
	ClientFingerprint string      `yaml:"client-fingerprint"`
}

type RealityOpts struct {
	PublicKey string `yaml:"public-key"`
	ShortId   string `yaml:"short-id"`
}
