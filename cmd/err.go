package cmd

import (
	"log"
)

type Err int

const (
	// OK��ʾ�ɹ� 
	OK Err = 0
	//ʱ���ͻ�������ڼ���ĳ������������������ʱ���ͻ 
	TimeConflict Err = 1
	// NoSuchMeeting��ʾ�û��鲻���� 
	NoSuchMeeting Err = 2
	// NoSuchUser ��ʾ���û������� 
	NoSuchUser Err = 3
	// DuplicateMeeting ��ʾҪ����һ���Ѿ����ڵĻ��� 
	DuplicateMeeting Err = 4
	// DuplicateUser��ʾ��ע��һ���Ѿ����ڵ��û� 
	DuplicateUser Err = 5
	// InvalidTime ��ʾʱ����ڴ��� 
	InvalidTime Err = 6
	// NoSuchFile ��ʾ����Ҫ�ҵ��ļ������� 
	NoSuchFile Err = 7
	// InconsistentState ��ʾ״̬���� 
	InconsistentState Err = 8
	// AuthenticateFail��ʾ�û������벻ƥ�� 
	AuthenticateFail Err = 9
	// WrongLoginState ��ʾ��¼ʱ״̬���� 
	WrongLoginState Err = 10
	// NotEnoughPrivilege ��ʾ������ͬ���Ĳ��� 
	NotEnoughPrivilege Err = 11
)

func LogFatalIfError(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
