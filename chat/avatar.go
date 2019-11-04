package main

import (
	"errors"
)

// ErrNoAvatarURL 는 Avatar 인스턴스가 아바타 URL을 제공할 수 없을 때 리턴되는 에러다.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

// Avatar 는 사용자 프로필 사진을 표현할 수 있는 타입을 나타낸다.
type Avatar interface {

	// GetAvatarURL 은 지정된 클라이언트에 대한 아바타URL을 가져오고, 문제가 발생하면 에러를 리턴한다.
	// 객체가 지정된 클라이언트의 URL을 가져올 수 없는 경우 ErrNoAvatarURL이 리턴된다.
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar 는 avatar 인터페이스를 구현한다.
type AuthAvatar struct{}

// UseAuthAvatar 는 ....
var UseAuthAvatar AuthAvatar

// GetAvatarURL 는 avatar 인터페이스의 GetAvatarURL 메서드를 구현한다.
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar 는 avatar 인터페이스를 구현한다.
type GravatarAvatar struct{}

// UseGravatar ...
var UseGravatar GravatarAvatar

// GetAvatarURL 는 GravatarAvatar 타입에 avatar 인터페이스 GetAvatarURL 메서드를 구현한다.
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
