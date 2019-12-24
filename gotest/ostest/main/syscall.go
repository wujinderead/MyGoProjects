package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

const EPOLLET = 0x80000000 // syscall.EPOLLET is -0x80000000, while net._EPOLLET is 0x80000000

var (
	ipv4lo       = [4]byte{127, 0, 0, 1}
	ipv4zero     = [4]byte{}
	ipv6lo       = [16]byte{15: 1}
	ipv6zero     = [16]byte{}
	v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}
)

func main() {
	//fileSyscall()
	//normalSyscall()
	//netSyscall()
	//flockSyscall()
	//epollSyscall()
	//ipv6Listen()
	//fmt.Println(getHostIpv4())
	//fmt.Println(getHostIpv6())
}

func fileSyscall() {
	// open file
	fd, err := syscall.Open("/tmp/testtmp", syscall.O_CREAT|syscall.O_WRONLY|syscall.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	fmt.Println("opened fd:", fd)

	// close file
	defer closeFd(fd)

	// get file stat
	stat := new(syscall.Stat_t)
	err = syscall.Fstat(fd, stat)
	if err != nil {
		fmt.Println("fstat err:", err)
		return
	}
	fmt.Println("size:", stat.Size)
	fmt.Println("b sz:", stat.Blksize)
	fmt.Println("blk :", stat.Blocks)
	fmt.Println("uid :", stat.Uid)
	fmt.Println("gid :", stat.Gid)
	fmt.Println("mode:", strconv.FormatUint(uint64(stat.Mode), 8))

	// write data
	n, err := syscall.Write(fd, []byte(`my name is van.
iam an artist, a performance artist.
i'm hired for people to fulfill their fantasies,
their DEEP ♂ DARK ♂ FANTASIES.
`))
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	fmt.Println("wrote:", n)

	// get size
	err = syscall.Fstat(fd, stat)
	if err != nil {
		fmt.Println("fstat err:", err)
		return
	}
	fmt.Println("size:", stat.Size)
	fmt.Println()

	// fd 0 is stdin
	// fd 1 is stdout
	// fd 2 is stderr
	n, err = syscall.Write(1, []byte("stdout\n")) // write to stdout
	fmt.Println(n, err)
	n, err = syscall.Write(2, []byte("stderr\n")) // write to stderr
	fmt.Println(n, err)
}

func normalSyscall() {
	// get time
	timee := new(syscall.Time_t)
	now, err := syscall.Time(timee)
	if err != nil {
		fmt.Println("get time err:", err)
		return
	}
	fmt.Println("now:", now, *timee)
	fmt.Println()

	// get working directory
	buf := make([]byte, syscall.PathMax)
	_p0 := unsafe.Pointer(&buf[0])
	r0, _, e1 := syscall.Syscall(syscall.SYS_GETCWD, uintptr(_p0), uintptr(len(buf)), 0)
	if e1 != 0 {
		fmt.Println("getcwd errno:", e1)
		return
	}
	fmt.Println("cwd:", string(buf[:r0])) // there is a '\0' in returned buf, because C use '\0\ to terminate a char*
	fmt.Println()

	// get system info
	sinfo := new(syscall.Sysinfo_t)
	err = syscall.Sysinfo(sinfo)
	if err != nil {
		fmt.Println("sys info err:", err)
		return
	}
	fmt.Println("sys Uptime   :", sinfo.Uptime)
	fmt.Println("sys Loads    :", sinfo.Loads)
	fmt.Println("sys Totalram :", sinfo.Totalram)
	fmt.Println("sys Freeram  :", sinfo.Freeram)
	fmt.Println("sys Sharedram:", sinfo.Sharedram)
	fmt.Println("sys Bufferram:", sinfo.Bufferram)
	fmt.Println("sys Totalswap:", sinfo.Totalswap)
	fmt.Println("sys Freeswap :", sinfo.Freeswap)
	fmt.Println("sys Procs    :", sinfo.Procs)
	fmt.Println()

	// get uts namespace info
	uname := new(syscall.Utsname)
	_, _, e1 = syscall.RawSyscall(syscall.SYS_UNAME, uintptr(unsafe.Pointer(uname)), 0, 0)
	if e1 > 0 {
		fmt.Println("uname errno:", e1)
		return
	}
	fmt.Println("uname Sysname   :", string((*(*[65]byte)(unsafe.Pointer(&uname.Sysname[0])))[:]))
	fmt.Println("uname Nodename  :", string((*(*[65]byte)(unsafe.Pointer(&uname.Nodename[0])))[:]))
	fmt.Println("uname Release   :", string((*(*[65]byte)(unsafe.Pointer(&uname.Release[0])))[:]))
	fmt.Println("uname Version   :", string((*(*[65]byte)(unsafe.Pointer(&uname.Version[0])))[:]))
	fmt.Println("uname Machine   :", string((*(*[65]byte)(unsafe.Pointer(&uname.Machine[0])))[:]))
	fmt.Println("uname Domainname:", string((*(*[65]byte)(unsafe.Pointer(&uname.Domainname[0])))[:]))
	fmt.Println()

	// file system info
	fs := new(syscall.Statfs_t)
	err = syscall.Statfs("/dev/sda", fs)
	if err != nil {
		fmt.Println("stat fs err:", err)
		return
	}
	fmt.Println("sda Type   :", strconv.FormatInt(fs.Type, 16))
	fmt.Println("sda Bsize  :", fs.Bsize)
	fmt.Println("sda Blocks :", fs.Blocks)
	fmt.Println("sda Bfree  :", fs.Bfree)
	fmt.Println("sda Bavail :", fs.Bavail)
	fmt.Println("sda Files  :", fs.Files)
	fmt.Println("sda Ffree  :", fs.Ffree)
	fmt.Println("sda Namelen:", fs.Namelen)
}

func flockSyscall() {
	ex := func(name string, latency time.Duration) {
		fd, err := syscall.Open("/tmp/testtmp", syscall.O_CREAT|syscall.O_WRONLY|syscall.O_APPEND, 0644)
		if err != nil {
			fmt.Println(name, "open err:", err)
			return
		}
		fmt.Println(name, "opened fd:", fd, time.Now())

		err = syscall.Flock(fd, syscall.LOCK_EX)
		if err != nil {
			fmt.Println(name, "exclusive lock err:", err)
			return
		}

		n, err := syscall.Write(fd, []byte(strconv.FormatInt(rand.Int63(), 3)+strconv.FormatInt(rand.Int63(), 3)))
		fmt.Println(name, "write:", n, time.Now())
		time.Sleep(latency)

		err = syscall.Flock(fd, syscall.LOCK_UN)
		if err != nil {
			fmt.Println(name, "exclusive unlock err:", err)
			return
		}

		closeFd(fd)
	}
	go ex("n1", 2*time.Second)
	time.Sleep(time.Second)
	go ex("n2", time.Second)
	time.Sleep(2 * time.Second)
}

func netSyscall() {
	serveraddr := &syscall.SockaddrInet4{Port: 12345, Addr: [4]byte{127, 0, 0, 1}}
	sockfd, err := createListenerSocket(serveraddr, false) // blocking socket
	if err != nil {
		fmt.Println("create listener socket err:", err)
		return
	}
	fmt.Println("create listener socket on addr:", serveraddr.Addr, ", port:", serveraddr.Port, ", sockfd:", sockfd)

	// set a dialer
	f := func(name, msg string, latency time.Duration) {
		time.Sleep(latency)
		// open a dialer socket
		sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
		if err != nil {
			fmt.Println(name, "socket err:", err)
			return
		}
		fmt.Println(name, "sockfd:", sockfd)

		// connect to remote
		err = syscall.Connect(sockfd, serveraddr)
		if err != nil {
			fmt.Println(name, "connect err:", err)
			return
		}

		// write some data
		for i := 0; i < 3; i++ {
			// write to socket, can also use 'send', 'sendto', sendmsg' with more options
			n, err := syscall.Write(sockfd, []byte(msg))
			if err != nil {
				fmt.Println(name, "write err:", err)
				return
			}
			fmt.Println(name, "wrote:", n)
			time.Sleep(time.Second / 2)
		}

		// close socket
		closeFd(sockfd)
		fmt.Println(name, "socket closed:", sockfd)
	}
	go f("dialer1", "my name is van.", time.Second)
	go f("dialer2", "change the boss of this gym.", 2*time.Second)

	// accept income connection
	for i := 0; i < 2; i++ {
		connfd, peer, err := syscall.Accept(sockfd)
		if err != nil {
			if err == syscall.EAGAIN { // socket set as non-blocking and no connections are present to be accepted
				continue
			}
			fmt.Println("accept err:", err)
			break
		}
		go func(connfd int, peer syscall.Sockaddr) {
			addr := peer.(*syscall.SockaddrInet4)
			fmt.Println("client dialer socket on addr:", addr.Addr, ", port:", addr.Port, ", connfd:", connfd)
			for {
				buf := make([]byte, 64)
				// read from socket, can also use 'recv', 'recvfrom', 'recvmsg' with more options
				n, err := syscall.Read(connfd, buf) // why non-blocking?
				if err != nil {
					fmt.Println("read err:", connfd, err)
					return
				}
				if n == 0 {
					time.Sleep(time.Second)
					continue
				}
				fmt.Printf("connfd %d reveive %s\n", connfd, string(buf[:n]))
			}
		}(connfd, peer)
	}
	time.Sleep(3 * time.Second)

	// close socket
	closeFd(sockfd)
	fmt.Printf("sockfd %d closed\n", sockfd)
}

func epollSyscall() {
	// create some listener
	hostv6, hostid := getHostIpv6()
	addrs := []syscall.Sockaddr{
		&syscall.SockaddrInet4{Port: 12344, Addr: ipv4lo},
		&syscall.SockaddrInet4{Port: 12345, Addr: getHostIpv4()},
		&syscall.SockaddrInet4{Port: 12346, Addr: ipv4zero},
		&syscall.SockaddrInet6{Port: 12347, Addr: ipv6lo, ZoneId: 1},
		&syscall.SockaddrInet6{Port: 12348, Addr: hostv6, ZoneId: uint32(hostid)},
		// &syscall.SockaddrInet6{Port: 12349, Addr: ipv6zero},   // can not dial ipv6 all 0 addr
	}
	sockfds := make([]int, 0)
	for i := range addrs {
		sockfd, err := createListenerSocket(addrs[i], true) // socket should be non-blocking for epoll
		if err != nil {
			fmt.Println("fail to listen on addr:", addrs[i], ", err:", err)
			continue
		}
		fmt.Println("listen on addr:", toString(addrs[i]), "sockfd:", sockfd)
		sockfds = append(sockfds, sockfd)
		defer closeFd(sockfd)
	}

	// create an epoller
	epfd, err := syscall.EpollCreate1(syscall.EPOLL_CLOEXEC)
	if err != nil {
		fmt.Println("epoll create err:", err)
		return
	}
	defer closeFd(epfd)

	// register listeners to the epoller. normally, it should be one-epoller-one-listener.
	// but epoll can handler multiple listeners for sure.
	for i := range sockfds {
		//set socket to non-blocking
		epollevent := &syscall.EpollEvent{
			Events: syscall.EPOLLIN | syscall.EPOLLRDHUP | syscall.EPOLLHUP | EPOLLET,
			Fd:     int32(sockfds[i]),
		}
		err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, sockfds[i], epollevent)
		if err != nil {
			fmt.Println("fail to epoll add sockfd:", sockfds[i])
		}
	}

	// get epoll events
	events := make([]syscall.EpollEvent, 10) // events buffer
	for {
		n, err := syscall.EpollWait(epfd, events, 10000) // max block 10s
		if err != nil {
			fmt.Println("epoll wait err:", err)
		}
		for j := 0; j < n; j++ {
			curfd := int(events[j].Fd)
			event := events[j].Events
			fmt.Println("event fd:", curfd, "type:", event)
			if event&syscall.EPOLLERR > 0 || event&syscall.EPOLLHUP > 0 || event&syscall.EPOLLRDHUP > 0 {
				fmt.Println("got event code:", event, ", fd:", curfd)
				closeFd(curfd) // got err, close fd
				continue
			}
			if curfd < sockfds[len(sockfds)-1] { // event fd == listener fd, accept new conn
				connfd, remoteaddr, err := syscall.Accept(curfd)
				if err != nil {
					fmt.Println("accept err for sockfd:", curfd)
					continue
				}
				fmt.Println("accept connfd:", connfd, "remote addr:", toString(remoteaddr))
				// set new conn non-blocking
				if setSocketNonBlocking(connfd) < 0 {
					fmt.Println("set nonblock err for connfd:", connfd)
					closeFd(connfd)
					continue
				}
				// register new conn to epoller
				newevent := new(syscall.EpollEvent)
				newevent.Fd = int32(connfd)
				newevent.Events = syscall.EPOLLIN | EPOLLET
				err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connfd, newevent)
				if err != nil {
					fmt.Println("fail to epoll add connfd:", connfd)
					closeFd(connfd)
					continue
				}
			} else { // readable event
				buf := make([]byte, 30)
				for {
					n, err := syscall.Read(curfd, buf) // handle read
					if n <= 0 || err == syscall.EAGAIN {
						break
					}
					fmt.Printf("read from connfd: %d, content: `%s`\n", curfd, string(buf[:n]))
					n, err = syscall.Write(curfd, buf[:n])
					fmt.Println("connfd:", curfd, ", write:", n, err)
				}
			}
		}
	}
}

func closeFd(fd int) {
	err := syscall.Close(fd)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "close fd %d err: %s\n", fd, err.Error())
	}
}

func createListenerSocket(addr syscall.Sockaddr, nonclock bool) (int, error) {
	// get sockfd
	sockfd := 0
	var err error
	flags := syscall.SOCK_STREAM | syscall.SOCK_CLOEXEC
	if nonclock {
		flags |= syscall.SOCK_NONBLOCK
	}
	switch addr.(type) {
	case *syscall.SockaddrInet4:
		sockfd, err = syscall.Socket(syscall.AF_INET, flags, 0)
		if err != nil {
			return -1, err
		}
	case *syscall.SockaddrInet6:
		sockfd, err = syscall.Socket(syscall.AF_INET6, flags, 0)
		if err != nil {
			return -1, err
		}
	default:
		return -1, errors.New("unsupported addr")
	}

	// bind to addr
	err = syscall.Bind(sockfd, addr)
	if err != nil {
		return -1, err
	}

	// set sockfd as a passive socket to accept connection
	err = syscall.Listen(sockfd, syscall.SOMAXCONN)
	if err != nil {
		return -1, err
	}
	return sockfd, nil
}

func getHostIpv4() [4]byte {
	hn, err := os.Hostname()
	if err != nil {
		fmt.Println("get hostname err:", err)
		return ipv4lo
	}
	ips, err := net.LookupIP(hn)
	if err != nil {
		fmt.Println("lookup ip err:", err)
		return ipv4lo
	}
	for i := range ips {
		if bytes.HasPrefix(ips[i], v4InV6Prefix) {
			return *(*[4]byte)(unsafe.Pointer(&ips[i].To4()[0]))
		}
	}
	return ipv4lo
}

func getHostIpv6() ([16]byte, int) {
	hn, err := os.Hostname()
	if err != nil {
		fmt.Println("get hostname err:", err)
		return ipv6lo, 1
	}
	ips, err := net.LookupIP(hn)
	if err != nil {
		fmt.Println("lookup ip err:", err)
		return ipv6lo, 1
	}
	hostv6 := []byte{}
	for i := range ips {
		if !bytes.HasPrefix(ips[i], v4InV6Prefix) {
			hostv6 = ips[i]
			break
		}
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get interfaces err:", err)
		return ipv6lo, 1
	}
	for i := range ifaces {
		addrs, err := ifaces[i].Addrs()
		if err != nil {
			fmt.Println("get interface addrs err:", err)
			continue
		}
		for j := range addrs {
			switch addrs[j].(type) {
			case *net.IPAddr:
				if bytes.Equal(addrs[j].(*net.IPAddr).IP, hostv6) {
					return *(*[16]byte)(unsafe.Pointer(&hostv6[0])), ifaces[i].Index
				}
			case *net.IPNet:
				if bytes.Equal(addrs[j].(*net.IPNet).IP, hostv6) {
					return *(*[16]byte)(unsafe.Pointer(&hostv6[0])), ifaces[i].Index
				}
			default:
				continue
			}
		}
	}
	return ipv6lo, 1
}

// test ipv6 listener
// ping6 -c 5 -I enp0s31f6 fe80::82fe:2beb:cc9f:530c
func ipv6Listen() {
	// listen on ethernet interface, [fe80::82fe:2beb:cc9f:530c%enp0s31f6]:12345
	// echo "msg" | nc -6N fe80::82fe:2beb:cc9f:530c%enp0s31f6 12345
	if false {
		addr := &net.TCPAddr{
			IP:   net.ParseIP("fe80::82fe:2beb:cc9f:530c"), // ipv6 addr
			Port: 12345,                                    // port
			Zone: "enp0s31f6",                              // the interface name where this ipv6 address is bound?
		}
		l, err := net.ListenTCP("tcp6", addr)
		if err != nil {
			fmt.Println("listen err:", err)
			return
		}
		fmt.Println(l.Addr().(*net.TCPAddr).Port)
		fmt.Println(l.Addr().(*net.TCPAddr).IP)
		fmt.Println(l.Addr().(*net.TCPAddr).Zone)
		for i := 0; i < 3; i++ {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("accept err:", err)
			}
			fmt.Println(conn.RemoteAddr().String()) // [fe80::aef2:c10b:1512:c191%enp0s31f6]:56964
			fmt.Println(conn.RemoteAddr().Network())
			fmt.Println("close:", conn.Close())
		}
	}

	// listen on ipv6 loopback
	// echo "msg" | nc -6N ::1 12345
	if false {
		lis, err := net.Listen("tcp6", "[::1]:12345") // [::1%lo]:12345 is also ok
		if err != nil {
			fmt.Println("listen err:", err)
			return
		}
		fmt.Println(lis.Addr().(*net.TCPAddr).Port)
		fmt.Println(lis.Addr().(*net.TCPAddr).IP)
		fmt.Println(lis.Addr().(*net.TCPAddr).Zone)
		for i := 0; i < 3; i++ {
			conn, err := lis.Accept()
			if err != nil {
				fmt.Println("accept err:", err)
			}
			fmt.Println("remote:", conn.RemoteAddr().String()) // [::1]:40530
			fmt.Println("network:", conn.RemoteAddr().Network())
			fmt.Println("local:", conn.LocalAddr().String()) // [::1]:12345
			fmt.Println("close:", conn.Close())
		}
	}

	// listen on ipv6 zero address
	// echo "msg" | nc -6N ::1 12345
	if true {
		lis, err := net.Listen("tcp6", "[::]:12345") // [::1%lo]:12345 is also ok
		if err != nil {
			fmt.Println("listen err:", err)
			return
		}
		fmt.Println(lis.Addr().(*net.TCPAddr).Port)
		fmt.Println(lis.Addr().(*net.TCPAddr).IP)
		fmt.Println(lis.Addr().(*net.TCPAddr).Zone)
		for i := 0; i < 3; i++ {
			conn, err := lis.Accept()
			if err != nil {
				fmt.Println("accept err:", err)
			}
			// local addr:                                   remote addr:
			//   [::1]:12345                                   [::1]:40568                                   // from localhost
			//   [fe80::82fe:2beb:cc9f:530c%enp0s31f6]:12345   [fe80::82fe:2beb:cc9f:530c%enp0s31f6]:52630   // from local ethernet
			//   [fe80::82fe:2beb:cc9f:530c%enp0s31f6]:12345   [fe80::aef2:c10b:1512:c191%enp0s31f6]:57716   // from remote ethernet
			fmt.Println("remote:", conn.RemoteAddr().String())
			fmt.Println("network:", conn.RemoteAddr().Network())
			fmt.Println("local:", conn.LocalAddr().String())
			fmt.Println("close:", conn.Close())
		}
	}
}

func toString(addr syscall.Sockaddr) string {
	switch addr.(type) {
	case *syscall.SockaddrInet4:
		ipv4 := addr.(*syscall.SockaddrInet4)
		return net.IPv4(ipv4.Addr[0], ipv4.Addr[1], ipv4.Addr[2], ipv4.Addr[3]).String() + ":" + strconv.Itoa(ipv4.Port)
	case *syscall.SockaddrInet6:
		ipv6 := addr.(*syscall.SockaddrInet6)
		return fmt.Sprintf("[%s%%%d]:%d", net.IP(ipv6.Addr[:]), ipv6.ZoneId, ipv6.Port)
	default:
		return ""
	}
}

func setSocketNonBlocking(sfd int) int {
	flags, _, errno := syscall.Syscall(syscall.SYS_FCNTL, uintptr(sfd), syscall.F_GETFL, 0)
	if errno != 0 {
		fmt.Println("fcntl get flags err:", errno)
		return -1
	}

	flags |= syscall.EPOLL_NONBLOCK
	_, _, errno = syscall.Syscall(syscall.SYS_FCNTL, uintptr(sfd), syscall.F_SETFL, flags)
	if errno != 0 {
		fmt.Println("fcntl set flags err:", errno)
		return -1
	}
	return 0
}
