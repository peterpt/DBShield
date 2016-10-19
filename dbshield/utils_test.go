package dbshield

import (
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/nim4/DBShield/dbshield/config"
	"github.com/nim4/mock"
)

func TestInitModel(t *testing.T) {
	config.Config.DBDir = os.TempDir()
	config.Config.DBType = "mysql"
	initModel()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic")
		}
	}()
	initModel()
}

func TestDbNameToStruct(t *testing.T) {
	_, err := dbNameToStruct("mysql")
	if err != nil {
		t.Error("Expected struct, got ", err)
		return
	}
	_, err = dbNameToStruct("oracle")
	if err != nil {
		t.Error("Expected struct, got ", err)
		return
	}
	_, err = dbNameToStruct("postgres")
	if err != nil {
		t.Error("Expected struct, got ", err)
		return
	}
	_, err = dbNameToStruct("invalid")
	if err == nil {
		t.Error("Expected error")
		return
	}
}

func TestInitLogging(t *testing.T) {
	config.Config.LogPath = "stdout"
	initLogging()
	config.Config.LogPath = "stderr"
	initLogging()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic")
		}
	}()
	config.Config.LogPath = "/tmp"
	initLogging()
}

func TestInitSignal(t *testing.T) {
	initSignal()
	time.Sleep(1 * time.Second)
}

func TestHandleClient(t *testing.T) {
	serverAddr, _ := net.ResolveTCPAddr("tcp", config.Config.TargetIP+":"+strconv.Itoa(int(config.Config.TargetPort)))
	var s mock.ConnMock
	err := handleClient(s, serverAddr)
	if err != nil {
		t.Error("Expected nil got error", err)
	}

	err = handleClient(s, nil)
	if err == nil {
		t.Error("Expected error got nil")
	}

	mock.ReturnError(true)
	defer mock.ReturnError(false)

	err = handleClient(s, serverAddr)
	if err == nil {
		t.Error("Expected error got nil")
	}
}
