package ipnetgen

import "net"

//Increment increments the given net.IP by one bit. Incrementing the last IP in an IP space (IPv4, IPV6) is undefined.
func Increment(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		//only add to the next byte if we overflowed
		if ip[i] != 0 {
			break
		}
	}
}

//IPNetGenerator is net.IPnet that you can iterate over
type IPNetGenerator struct {
	*net.IPNet
	count uint64

	//state
	idx     uint64
	current net.IP
}

//NewIPNetGenerator creates a new IPNetGenerator from a CIDR string, or an error if the CIDR is invalid.
func NewIPNetGenerator(cidr string) (*IPNetGenerator, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	ones, bits := ipNet.Mask.Size()

	newIP := make(net.IP, len(ipNet.IP))
	copy(newIP, ipNet.IP)

	return &IPNetGenerator{
		IPNet:   ipNet,
		count:   1 << uint8(bits-ones),
		current: newIP,
	}, nil
}

//Next returns the next net.IP in the subnet
func (g *IPNetGenerator) Next() net.IP {
	if g.idx > g.count-1 {
		return nil
	}
	current := make(net.IP, len(g.current))
	copy(current, g.current)
	Increment(g.current)
	g.idx++

	return current
}
