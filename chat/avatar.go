package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatarURL 는 Avatar 인스턴스가 아바타 URL을 제공할 수 없을 때 리턴되는 에러다.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// Avatar 는 사용자 프로필 사진을 표현할 수 있는 타입을 나타낸다.
type Avatar interface {

	// GetAvatarURL 은 지정된 클라이언트에 대한 아바타URL을 가져오고, 문제가 발생하면 에러를 리턴한다.
	// 객체가 지정된 클라이언트의 URL을 가져올 수 없는 경우 ErrNoAvatarURL이 리턴된다.
	GetAvatarURL(ChatUser) (string, error)
}

// TryAvatars ...
type TryAvatars []Avatar

// GetAvatarURL ...
func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

// AuthAvatar 는 avatar 인터페이스를 구현한다.
type AuthAvatar struct{}

// GetAvatarURL 는 avatar 인터페이스의 GetAvatarURL 메서드를 구현한다.
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

// UseAuthAvatar 는 ....
var UseAuthAvatar AuthAvatar

// GravatarAvatar 는 avatar 인터페이스를 구현한다.
type GravatarAvatar struct{}

// GetAvatarURL 는 GravatarAvatar 타입에 avatar 인터페이스 GetAvatarURL 메서드를 구현한다.
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

// UseGravatar ...
var UseGravatar GravatarAvatar

// FileSystemAvatar ...
type FileSystemAvatar struct{}

// GetAvatarURL ...
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}

// UseFileSystemAvatar ...
var UseFileSystemAvatar FileSystemAvatar
