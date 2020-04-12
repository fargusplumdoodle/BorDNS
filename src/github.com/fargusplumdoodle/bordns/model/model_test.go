package model

import (
	"context"
	"github.com/fargusplumdoodle/bordns/conf"
	"testing"
)

const (
	ETCD_HOST = "172.19.0.3:2379"
)

func TestGetZoneFromHostNotFound(t *testing.T) {
	/*
		Checks to see if hosts were found properly
	*/
	// setting zones
	conf.Env = &conf.Config{
		EtcdHosts:  nil,
		ListenAddr: "",
		Zones: []conf.ZoneConfig{
			{Zone: "ey"},
			{Zone: "wont.match.bor"},
			{Zone: "should.match.test.bor"},
		},
	}

	// making table
	notFoundZones := []string{
		"this.should.match.test.bor.oops.no.it.wont",
		"hey.ey.ey.not.today.amirite",
	}

	for _, x := range notFoundZones {
		result, err := getZoneFromHost(x)
		if err == nil {
			t.Errorf("expected result to be not found, instead got %q", result)
		}
	}
}
func TestGetZoneFromHostFound(t *testing.T) {
	/*
		Checks to see if hosts were found properly
	*/
	// setting zones
	conf.Env = &conf.Config{
		EtcdHosts:  nil,
		ListenAddr: "",
		Zones: []conf.ZoneConfig{
			{Zone: "ey"},
			{Zone: "wont.match.bor"},
			{Zone: "should.match.test.bor"},
		},
	}

	// making table
	table := []struct {
		in  string
		out conf.ZoneConfig
	}{
		{"this.should.match.test.bor", conf.ZoneConfig{Zone: "should.match.test.bor"}},
		{"hey.ey.ey", conf.ZoneConfig{Zone: "ey"}},
	}

	for _, x := range table {
		result, err := getZoneFromHost(x.in)

		if err != nil {
			t.Errorf("expected: %q for %q", x.out.Zone, err.Error())

		} else if result.Zone != x.out.Zone {
			t.Errorf("expected: %q for %q", x.out.Zone, result.Zone)
		}
	}
}
func TestGetReversedDomain(t *testing.T) {
	// ensures that domain names are properly converted to slices and reversed
	table := []struct {
		in  string
		out []string
	}{
		{"test.bor", []string{"bor", "test"}},
		{"longer.test.bor", []string{"bor", "test", "longer"}},
	}

	for _, x := range table {
		result := getReversedDomain(x.in)

		if len(result) != len(x.out) {
			t.Errorf("expected: %q for %q", x.out, result)
		}

		for i := 0; i < len(x.out); i++ {
			if result[i] != x.out[i] {
				t.Errorf("expected: %q for %q", x.out, result)
			}
		}

	}

}
func TestGetUnReversedDomain(t *testing.T) {
	// ensures that domain names are properly converted to back into
	// their origio
	table := []struct {
		out string
		in  []string
	}{
		{"test.bor", []string{"bor", "test"}},
		{"longer.test.bor", []string{"bor", "test", "longer"}},
	}

	for _, x := range table {
		result := unReverseDomain(x.in)

		if result != x.out {
			t.Errorf("expected: %q for %q", x.out, result)
		}
	}

}

/*
Test Add A Record
------------------
1. Set zones and create test table
2. Call AddARecord for each test in the table
3. Ensure IP was set to the appropriate path in Etcd
*/
func TestAddARecordValid(t *testing.T) {
	// 1.  setting zones and making table
	conf.Env = &conf.Config{
		EtcdHosts:  []string{ETCD_HOST},
		ListenAddr: "",
		Zones: []conf.ZoneConfig{
			{Zone: "bor", EtcdPath: "/bor"},
			{Zone: "sekhnet.ra", EtcdPath: "/ra/sekhnet"},
		},
	}
	SetupDB(conf.Env.EtcdHosts)
	table := []struct {
		in  []string // 0. Host, 1. IP
		out string   // out is the expected etcd path.
	}{
		{[]string{"test.bor", "10.0.0.1"}, "/bor/bor/test"},
		{[]string{"something.plumdoodle.sekhnet.ra", "10.0.1.1"}, "/ra/sekhnet/ra/sekhnet/plumdoodle/something"},
		{[]string{"another.test.bor", "10.0.2.1"}, "/bor/bor/test/another"},
	}

	// 2.
	for _, x := range table {
		err := AddARecord(x.in[0], x.in[1]) // 0. Host, 1. IP

		if err != nil {
			t.Errorf("unexpected error: %q", err)
		}

		ctx, cancel := context.WithTimeout(context.TODO(), conf.DB_TIMEOUT)
		resp, err := client.Get(ctx, x.out)
		cancel()

		if err != nil {
			t.Fatalf("Error occured retrieving value of %q:  %q", x.in[0], err)
		}
		if resp == nil {
			t.Fatal("nil response!!")
			return
		}
		var passed = false
		for _, ev := range resp.Kvs {
			if string(ev.Value) != getARecordValue(x.in[1]) {
				t.Fatalf("invalid value in etcd: got %q, expected %q", ev.Value, getARecordValue(x.in[1]))
			} else {
				passed = true
				break
			}
		}
		if !passed {
			t.Fatalf("did not find values in database. host: %q  expected: %q", x.in[0], getARecordValue(x.in[1]))
		}
	}

}

// TODO: INVALID
