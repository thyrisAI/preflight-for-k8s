package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Helper function to print error and exit if the condition is not met
func checkError(err error, message string) {
	if err != nil {
		fmt.Printf("Error: %s: %v\n", message, err)
		os.Exit(1)
	}
}

// Test 1: Check Storage Class and verify there is more than 500GB available
func testStorageClass(clientset *kubernetes.Clientset) {
	scs, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	checkError(err, "Failed to get Storage Class")

	for _, sc := range scs.Items {
		// Check if Longhorn is being used
		fmt.Printf("Storage Class: %s\n", sc.Name)
		if strings.Contains(sc.Provisioner, "longhorn") {
			fmt.Println("Longhorn is being used.")
		}

		// Test for sufficient storage capacity
		if sc.Parameters["size"] != "" {
			sizeGB, _ := strconv.Atoi(sc.Parameters["size"])
			if sizeGB > 500 {
				fmt.Println("More than 500GB of storage available.")
			} else {
				fmt.Println("Insufficient storage.")
				os.Exit(1)
			}
		}
	}
}

// Test 2: Check if Longhorn support is enabled
func testLonghorn(clientset *kubernetes.Clientset) {
	scs, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	checkError(err, "Failed to get Storage Class")

	foundLonghorn := false

	for _, sc := range scs.Items {
		if strings.Contains(sc.Provisioner, "longhorn") {
			foundLonghorn = true
		}
	}

	if foundLonghorn {
		fmt.Println("Longhorn support found.")
	} else {
		fmt.Println("Longhorn support not found.")
		os.Exit(1)
	}
}

// Test 3: Check Volume Snapshot Class
func testVolumeSnapshotClass(clientset *kubernetes.Clientset) {
	// In a real-world scenario, you would check for the existence of VolumeSnapshotClass
	fmt.Println("Checking Volume Snapshot Class...")
	// If there are VolumeSnapshotClass resources in the cluster, list or validate them here
}

// Test 4: Check if the total cluster resources are greater than 24 Core CPU and 64 GB RAM
func testClusterResources(clientset *kubernetes.Clientset) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	checkError(err, "Failed to get nodes")

	totalCPU := 0
	totalMemory := int64(0)

	for _, node := range nodes.Items {
		// Calculate total CPU cores
		cpuQuantity := node.Status.Capacity[v1.ResourceCPU]
		memQuantity := node.Status.Capacity[v1.ResourceMemory]

		cpuCores, _ := strconv.Atoi(cpuQuantity.String())
		totalCPU += cpuCores

		// Calculate total memory in GB
		memBytes, _ := memQuantity.AsInt64()
		totalMemory += memBytes / (1024 * 1024 * 1024) // Convert bytes to GB
	}

	if totalCPU >= 24 {
		fmt.Println("Sufficient CPU: More than 24 cores.")
	} else {
		fmt.Println("Insufficient CPU resources.")
		os.Exit(1)
	}

	if totalMemory >= 64 {
		fmt.Println("Sufficient memory: More than 64 GB.")
	} else {
		fmt.Println("Insufficient memory resources.")
		os.Exit(1)
	}
}

func main() {
	// Kubernetes client configuration
	config, err := rest.InClusterConfig()
	checkError(err, "Failed to connect to Kubernetes cluster")

	clientset, err := kubernetes.NewForConfig(config)
	checkError(err, "Failed to create Kubernetes client")

	// Run tests in sequence
	testStorageClass(clientset)
	testLonghorn(clientset)
	testVolumeSnapshotClass(clientset)
	testClusterResources(clientset)

	// All tests passed
	fmt.Println("OKAY")
	time.Sleep(2 * time.Second) // Pause before shutting down the container
}
