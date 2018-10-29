package main

import "fmt"

/*
2017-08-18 10:31:26,165 INFO  org.apache.flink.runtime.io.network.netty.NettyConfig         - NettyConfig [server address: taskmanager2/172.25.100.5, server port: 0, ssl enabled: false, memory segment size (bytes): 32768, transport type: NIO, number of server threads: 4 (manual), number of client threads: 4 (manual), server connect backlog: 0 (use Netty's default), client connect timeout (sec): 120, send/receive buffer size (bytes): 0 (use Netty's default)]
2017-08-18 10:31:26,174 INFO  org.apache.flink.runtime.taskexecutor.TaskManagerConfiguration  - Messages have a max timeout of 10000 ms
2017-08-18 10:31:26,177 INFO  org.apache.flink.runtime.taskexecutor.TaskManagerServices     - Temporary file directory '/tmp': total 437 GB, usable 364 GB (83.30% usable)
2017-08-18 10:31:26,213 INFO  org.apache.flink.runtime.io.network.buffer.NetworkBufferPool  - Allocated 101 MB for network buffer pool (number of memory segments: 3252, bytes per segment: 32768).
2017-08-18 10:31:26,340 INFO  org.apache.flink.runtime.io.network.NetworkEnvironment        - Starting the network environment and its components.
2017-08-18 10:31:26,345 INFO  org.apache.flink.runtime.io.network.netty.NettyClient         - Successful initialization (took 2 ms).
2017-08-18 10:31:26,366 INFO  org.apache.flink.runtime.io.network.netty.NettyServer         - Successful initialization (took 21 ms). Listening on SocketAddress /172.25.100.5:33544.

*/

func pushDominoes(dominoes string) string {
	var ret = []byte(dominoes)
	i := 0
	R := -1
	L := -1
	for i < len(dominoes) {
		ret[i] = dominoes[i]
		if dominoes[i] == 'R' {
			if R > L {
				for j := R; j < i; j++ {
					ret[j] = 'R'
				}
			}
			R = i
		}
		if dominoes[i] == 'L' {
			if L > R || R == -1 {
				for j := L + 1; j < i; j++ {
					ret[j] = 'L'
				}
			} else if R > L {
				r := R + 1
				l := i - 1
				for r < l {
					ret[r] = 'R'
					ret[l] = 'L'
					r++
					l--
				}
			}
			L = i
		}
		i++
	}
	if R > L {
		for i := R; i < len(dominoes); i++ {
			ret[i] = 'R'
		}
	}
	return string(ret)
}

func main() {
	in := []string{".L.L.R.R...L.L.R..L..R..R..", ".L.R...LR..L..", "RR.L"}
	for _, str := range in {
		fmt.Println(str)
		fmt.Println(pushDominoes(str))
	}
}
