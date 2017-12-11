package messageBuilder

import (
	"testing"
	"serverList"
	"cryptographer"
	"message"
	"fmt"
	"bytes"
	"conversationMessageBuilder"
	"commonKeyProtocol"
	"receiverKeyHandler"
)

func TestMessageBuilder_Build(t *testing.T) {
	msgb := New(serverList.NewServerList())
	msgb.SetMyName("10000000").SetMyServer("00000000")

	expected := ([]byte)("  00000002  00000001  recvserv  recvrecv  0000000010000000  abcd")
	expected[0] = 0
	expected[1] = 0
	expected[10] = 0
	expected[11] = 0
	expected[20] = 0
	expected[21] = 0
	expected[30] = 0
	expected[31] = 0
	expected[40] = 0
	expected[41] = 0
	expected[58] = 0
	expected[59] = 0

	msgString := "abcd"

	path := []string{"00000001", "00000002"}

	convoBuilder := conversationMessageBuilder.NewConversationMessageBuilder(&commonKeyProtocol.CommonKeyProtocolImpl{})

	msg,_ :=
		msgb.SetMsgString(msgString).
		SetMsgContentBuilder(convoBuilder).
		SetReceiverKeyHandler(&receiverKeyHandler.ReceiverKeyHandlerImpl{}).
		SetReceiver("recvrecv").
		SetReceiverServer("recvserv").
		SetEncType(cryptographer.PLAIN).
		SetMsgType(message.DEFAULT).
		SetPath(path).
		Build()
	fmt.Println(string(msg.ToBytes()))
	fmt.Println(string(expected))

	if !bytes.Equal(msg.ToBytes(), ([]byte)(expected)){
		t.Error("Unexpected message")
	}
}
