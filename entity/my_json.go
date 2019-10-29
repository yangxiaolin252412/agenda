package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func User_JsonDecode(js []byte) User {
	var jm User
	err := json.Unmarshal(js, &jm)
	if err != nil {
		fmt.Println("error2")
	}
	return jm
}

func User_JsonEncode(m User) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error1")
		os.Exit(1)
	}
	return data
}

func User_ReadFromFile() []User {
	var tmp []User
	f, err := os.Open("entity/data/User.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		tmp = append(tmp, User_JsonDecode([]byte(line)))
	}
	return tmp
}

func User_WriteToFile(My_User []User) {
	file, err := os.OpenFile("entity/data/User.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/User.txt", 0)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(My_User); i++ {
		file.WriteString(string(User_JsonEncode(My_User[i])[:]))
		file.WriteString("\n")
	}
}

func Meeting_JsonDecode(js []byte) Meeting {
	var jm Meeting
	err := json.Unmarshal(js, &jm)
	if err != nil {
		fmt.Println("error2")
	}
	return jm
}

func Meeting_JsonEncode(m Meeting) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error1")
		os.Exit(1)
	}
	return data
}

func Meeting_ReadFromFile() []Meeting {
	var tmp []Meeting
	f, err := os.Open("entity/data/Meeting.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		tmp = append(tmp, Meeting_JsonDecode([]byte(line)))
	}
	return tmp
}

func Meeting_WriteToFile(My_Meeting []Meeting) {
	file, err := os.OpenFile("entity/data/Meeting.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/Meeting.txt", 0)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(My_Meeting); i++ {
		file.WriteString(string(Meeting_JsonEncode(My_Meeting[i])[:]))
		file.WriteString("\n")
	}
}

func LN_ReadFromFile() []string {
	var tmp []string
	f, err := os.Open("entity/data/Host.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		tmp = append(tmp, line)
	}
	return tmp
}

func LN_WriteToFile(name string) {
	file, err := os.OpenFile("entity/data/Host.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/Host.txt", 0)
	if err != nil {
		fmt.Println("open file failed.", err.Error())
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(name)
	file.WriteString("\n")
}

func Empty_login() {
	os.Truncate("entity/data/Host.txt", 0)
}
