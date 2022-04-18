package server

import (
	"fmt"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"log"
)


func StartServer()  {
	fmt.Printf("cuda nvml start\n")
	ret := nvml.Init()
	if ret != nvml.SUCCESS{
		log.Fatalf("unable to init NVML: %v\n", nvml.ErrorString(ret))
	}
	defer func() {
		ret:= nvml.Shutdown()
		if ret != nvml.SUCCESS{
			log.Fatalf("unable to shutdown NVML: %v\n", nvml.ErrorString(ret))
		}
	}()

	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS{
		log.Fatalf("unable to get device count: %v\n", nvml.ErrorString(ret))
	}

	for i:=0;i<count;i++{
		device, ret := nvml.DeviceGetHandleByIndex(i)
		if ret != nvml.SUCCESS{
			log.Fatalf("unable to get device at index %v: %v\n", i, nvml.ErrorString(ret))
		}
		uuid, ret := device.GetUUID()
		if ret!= nvml.SUCCESS{
			log.Fatalf("unable to get UUID Of device at index %v: %v\n", i, nvml.ErrorString(ret))
		}
		fmt.Printf("UUID: %v\n", uuid)
	}
}
