package main

import "fmt"

type SocialMediaType string

const (
	Facebook  SocialMediaType = "Facebook"
	Instagram SocialMediaType = "Instagram"
	Twitter   SocialMediaType = "Twitter"
)

type SocialMediaInterface interface {
	getName() string
	getURL() string
	getType() SocialMediaType
}

type SocialMedia struct {
	Name string
	URL  string
	Type SocialMediaType
}

func (s *SocialMedia) getName() string {
	return s.Name
}

func (s *SocialMedia) getURL() string {
	return s.URL
}

func (s *SocialMedia) getType() SocialMediaType {
	return s.Type
}

func socialMedia(mediaType SocialMediaType) SocialMediaInterface {
	switch mediaType {
	case Facebook:
		return &SocialMedia{
			Name: "FACEBOOK",
			URL:  "https://www.facebook.com",
			Type: Facebook,
		}
	case Instagram:
		return &SocialMedia{
			Name: "INSTAGRAM",
			URL:  "https://www.instagram.com",
			Type: Instagram,
		}
	case Twitter:
		return &SocialMedia{
			Name: "TWITTER",
			URL:  "https://www.twitter.com",
			Type: Twitter,
		}
	}
	return nil
}

func main() {
	sosmedA := socialMedia(Facebook)
	_ = socialMedia(Twitter)
	_ = socialMedia(Instagram)

	fmt.Println(sosmedA.getName())
	fmt.Println(sosmedA.getURL())
	fmt.Println(sosmedA.getType())
}
