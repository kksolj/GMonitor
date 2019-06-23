package main

import (
	"encoding/json"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func Service() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		tk := req.Header.Get("TOKEN")
		if tk != Conf.Server.Token {
			writer.WriteHeader(401)
			return
		}
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(200)
		writer.Write(tempMonitor())
	})
	log.Printf("start server on %v \n", Conf.Server.Addr)
	if e := http.ListenAndServe(Conf.Server.Addr, nil); e != nil {
		log.Fatal("start server error", e)
	} else {

	}
}

type Status struct {
	Host        host.InfoStat             `json:"host"`
	Temperature []host.TemperatureStat    `json:"temperature"`
	User        []host.UserStat           `json:"user"`
	Cpu         float64                   `json:"cpu"`
	Memory      mem.VirtualMemoryStat     `json:"memory"`
	Swap        mem.SwapMemoryStat        `json:"swap"`
	Disk        []disk.UsageStat          `json:"disk"`
	Process     []*process.Process        `json:"process"`
	Docker      []docker.CgroupDockerStat `json:"docker"`
	Time        int64                     `json:"time"`
}

func tempMonitor() []byte {
	r := new(Status)
	i, _ := host.Info()
	r.Host = *i
	t, _ := host.SensorsTemperatures()
	r.Temperature = t
	u, _ := host.Users()
	r.User = u

	path, _ := disk.Partitions(true)
	if r.Disk == nil {
		r.Disk = make([]disk.UsageStat, len(path))
	}
	for _, x := range path {
		s, _ := disk.Usage(x.Mountpoint)
		r.Disk = append(r.Disk, *s)
	}
	v, _ := mem.VirtualMemory()
	r.Memory = *v
	s, _ := mem.SwapMemory()
	r.Swap = *s
	r.Cpu = cpuTime()
	p, _ := process.Processes()
	for _, ps := range p {
		ps.MemoryInfo()
		ps.CPUPercent()
	}
	r.Process = p
	_, err := exec.LookPath("docker")
	if err == nil {
		dc, _ := docker.GetDockerStat()
		r.Docker = dc
	}
	r.Time = time.Now().Unix()
	data, _ := json.Marshal(r)
	return data
}

/*func tempMonitor(dur time.Duration, close chan bool) <-chan string {
	out := make(chan string)
	path, _ := disk.Partitions(true)
	go func(d time.Duration, c <-chan bool, o chan<- string) {
	l:
		for {
			select {
			case <-c:
				break l
			default:
				buf := new(bytes.Buffer)
				diskStat := make([]string, 0, len(path))
				for _, x := range path {
					s, _ := disk.Usage(x.Mountpoint)
					diskStat = append(diskStat, fmt.Sprintf(`{"path":"%v","total":%v,"used":%v,"used_percent":%.2f}`, s.Path, s.Total, s.Used, s.UsedPercent))
				}
				dc, _ := docker.GetDockerStat()
				dockerStat := make([]string, 0, len(dc))
				for _, t := range dc {
					dockerStat = append(dockerStat, fmt.Sprintf(`{"conainter":"%v","containerId":"%v","status":%v}`, t.Name, t.ContainerID, t.Running))
				}
				time.Sleep(d)
			}
		}
	}(dur, close, out)
	return out
}*/
