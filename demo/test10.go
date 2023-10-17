package main

func main() {

	a := make([]string, 0)
	bb := []struct {
		groupName string
		name      string
	}{
		{
			groupName: "11",
			name:      "1",
		},
		{
			groupName: "22",
			name:      "2",
		},
		{
			groupName: "11",
			name:      "3",
		},
	}
	bol := false
	for _, v := range bb {
		for _, vv := range a {
			if vv == v.groupName {
				bol = false
			}
		}
		if bol {
			a = append(a, v.groupName)
		}

	}
}
