package main

import "fmt"

// 工单部门角色转换
type userStatus struct {
	deptID []int // 记录的部门的ID
	isLD   []int // 记录的是是否是部门主管和deptID对应，1为主管，0为不是主管
}

func getUserRole(status userStatus, saleDeptIDs []int) (map[int]string, map[int]bool) {
	roles := make(map[int]string)   // 部门ID对应的角色
	isManager := make(map[int]bool) // 部门ID对应的是否是主管

	// 判断用户角色
	for i := 0; i < len(status.deptID); i++ {
		deptID := status.deptID[i]
		ld := status.isLD[i]

		if contains(saleDeptIDs, deptID) {
			roles[deptID] = "销售"
		} else if ld == 1 {
			roles[deptID] = "主管"
		} else {
			roles[deptID] = "讲师"
		}

		isManager[deptID] = (ld == 1)
	}

	return roles, isManager
}

func contains(arr []int, num int) bool {
	for _, n := range arr {
		if n == num {
			return true
		}
	}
	return false
}

func main() {
	status := userStatus{
		deptID: []int{30, 31},
		isLD:   []int{0, 1},
	}
	saleDeptIDs := []int{30, 31, 32, 46, 2, 8, 9, 24, 43}

	roles, isManager := getUserRole(status, saleDeptIDs)

	fmt.Println("部门角色:")
	for deptID, role := range roles {
		fmt.Printf("部门ID: %d, 角色: %s\n", deptID, role)
	}

	fmt.Println("是否是主管:")
	for deptID, manager := range isManager {
		fmt.Printf("部门ID: %d, 是否是主管: %t\n", deptID, manager)
	}
}
