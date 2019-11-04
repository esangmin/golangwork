package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	// 값 설정
	testURL := "http://url-to-gravatar/"
	client.userData = map[string]interface{}{"avatar_url": testURL}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	}
	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar

	client := new(client)
	client.userData = map[string]interface{}{"userid": "2d7c5cde1400fe8d1d84444840590c5b"}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}

	if url != "//www.gravatar.com/avatar/2d7c5cde1400fe8d1d84444840590c5b" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}