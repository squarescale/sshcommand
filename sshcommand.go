package sshcommand

import (
	"github.com/docopt/docopt-go"
	"strings"
)

type SshCommand struct {
	argv []string
	args docopt.Opts
}

func (c *SshCommand) Argv() []string {
	return c.argv
}

func (c *SshCommand) Hostname() string {
	d := c.args["DESTINATION"].(string)
	if strings.Contains(d, "@") {
		d = strings.Split(d, "@")[1]
	}
	return d
}

func PrependOpt(argv []string, opt []string) []string {
	margv := append([]string{}, argv[0])
	margv = append(margv, opt...)
	margv = append(margv, argv[1:]...)
	return margv
}

func New(argv []string) (*SshCommand, error) {
	const ssh_usage = `
        usage: ssh [-46AaCfGgKkNnqsTVXxYy] [-v...] [-M...] [-t...] [-B <bind_interface>]
   [-b <bind_address>] [-c <cipher_spec>] [-D <dynamic>]
   [-E <log_file>] [-e <escape_char>] [-F <configfile>] [-I <pkcs11>]
   [-i <identity_file>...] [-J <jumpspec>] [-L <address>...]
   [-l <login_name>] [-m <mac_spec>] [-O <ctl_cmd>] [-o <option>] [-p <port>]
   [-Q <query_option>] [-R <address>...] [-S <ctl_path>] [-W <host:port>]
   [-w <tunspec>] DESTINATION [COMMAND...]
options:
    -B <bind_interface>
    -b <bind_address>
    -c <cipher_spec>
    -D <dynamic>
    -E <log_file>
    -e <espace_char>
    -F <configfile>
    -I <pkcs11>
    -i <identity_file>...
    -J <jumpspec>
    -L <address>...
    -l <login_name>
    -m <mac_spec>
    -O <ctl_cmd>
    -o <option>
    -p <port>
    -Q <query_option>
    -R <address>...
    -S <ctl_path>
    -W <host_port>
    -w <tunspec>
`
	parser := &docopt.Parser{
		HelpHandler:   docopt.NoHelpHandler,
		SkipHelpFlags: true,
	}

	sc := &SshCommand{}

	opts, err := parser.ParseArgs(ssh_usage, argv[1:], "")
	if err != nil {
		return sc, err
	}
	sc.argv = argv
	sc.args = opts
	return sc, nil
}
