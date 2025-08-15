package ifs

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"github.com/saichler/l8types/go/types"
)

const (
	Deleted_Entry = "__DD__"
)

type IElements interface {
	Elements() []interface{}
	Keys() []interface{}
	Errors() []error
	Element() interface{}
	Query(IResources) (IQuery, error)
	Key() interface{}
	Error() error
	Serialize() ([]byte, error)
	Deserialize([]byte, IRegistry) error
	Notification() bool
	ReplicasRequest() bool
	Append(IElements)
	AsList(IRegistry) (interface{}, error)
}

type IQuery interface {
	RootType() *types.RNode
	Properties() []IProperty
	Criteria() IExpression
	KeyOf() string
	Match(interface{}) bool
}

type IProperty interface {
	PropertyId() (string, error)
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{}) (interface{}, interface{}, error)
	Node() *types.RNode
	Parent() IProperty
	IsString() bool
	Resources() IResources
}

type IExpression interface {
	Condition() ICondition
	Operator() string
	Next() IExpression
	Child() IExpression
}

type ICondition interface {
	Comparator() IComparator
	Operator() string
	Next() ICondition
}

type IComparator interface {
	Left() string
	LeftProperty() IProperty
	Right() string
	RightProperty() IProperty
	Operator() string
}

func FormatTop(top *types.Top) string {
	if top == nil || len(top.Healths) == 0 {
		return "No processes running"
	}
	
	var sb strings.Builder
	
	currentTime := time.Now().Format("15:04:05")
	sb.WriteString(fmt.Sprintf("top - %s up 0 days, %d users,  load average: 0.00, 0.00, 0.00\n",
		currentTime, len(top.Healths)))
	
	totalTasks := len(top.Healths)
	running := 0
	sleeping := 0
	stopped := 0
	for _, health := range top.Healths {
		switch health.Status {
		case types.HealthState_Up:
			running++
		case types.HealthState_Down:
			stopped++
		default:
			sleeping++
		}
	}
	
	sb.WriteString(fmt.Sprintf("Tasks: %d total, %d running, %d sleeping, %d stopped, 0 zombie\n",
		totalTasks, running, sleeping, stopped))
	
	var totalCpu float64
	var totalMem uint64
	for _, health := range top.Healths {
		if health.Stats != nil {
			totalCpu += health.Stats.CpuUsage
			totalMem += health.Stats.MemoryUsage
		}
	}
	
	sb.WriteString(fmt.Sprintf("%%Cpu(s): %5.1f us,  0.0 sy,  0.0 ni, %5.1f id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st\n",
		totalCpu, 100.0-totalCpu))
	sb.WriteString(fmt.Sprintf("MiB Mem : %8.1f total,     0.0 free, %8.1f used,     0.0 buff/cache\n",
		float64(totalMem)/1024/1024, float64(totalMem)/1024/1024))
	sb.WriteString("MiB Swap:      0.0 total,      0.0 free,      0.0 used.      0.0 avail Mem\n")
	sb.WriteString("\n")
	
	sb.WriteString("  PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND\n")
	
	type processInfo struct {
		pid     string
		user    string
		virt    uint64
		res     uint64
		shr     uint64
		status  string
		cpu     float64
		mem     float64
		time    string
		command string
	}
	
	var processes []processInfo
	for key, health := range top.Healths {
		pid := strings.Split(key, "-")[0]
		if len(pid) > 8 {
			pid = pid[:8]
		}
		
		user := "root"
		if len(health.Alias) > 8 {
			user = health.Alias[:8]
		} else if health.Alias != "" {
			user = health.Alias
		}
		
		var virt, res, shr uint64
		var cpu, mem float64
		if health.Stats != nil {
			virt = health.Stats.MemoryUsage
			res = health.Stats.MemoryUsage
			shr = health.Stats.MemoryUsage / 4
			cpu = health.Stats.CpuUsage
			mem = cpu / 10
		}
		
		var status string
		switch health.Status {
		case types.HealthState_Up:
			status = "R"
		case types.HealthState_Down:
			status = "T"
		default:
			status = "S"
		}
		
		uptime := time.Duration(0)
		if health.StartTime > 0 {
			uptime = time.Since(time.Unix(health.StartTime, 0))
		}
		timeStr := fmt.Sprintf("%02d:%02d", int(uptime.Minutes()), int(uptime.Seconds())%60)
		
		command := health.Alias
		if command == "" {
			command = key
		}
		if len(command) > 15 {
			command = command[:15]
		}
		
		processes = append(processes, processInfo{
			pid:     pid,
			user:    user,
			virt:    virt,
			res:     res,
			shr:     shr,
			status:  status,
			cpu:     cpu,
			mem:     mem,
			time:    timeStr,
			command: command,
		})
	}
	
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].cpu > processes[j].cpu
	})
	
	for _, proc := range processes {
		sb.WriteString(fmt.Sprintf("%5s %-8s  20   0 %7d %7d %7d %s %5.1f %5.1f %8s %s\n",
			proc.pid,
			proc.user,
			proc.virt/1024,
			proc.res/1024,
			proc.shr/1024,
			proc.status,
			proc.cpu,
			proc.mem,
			proc.time,
			proc.command))
	}
	
	return sb.String()
}
