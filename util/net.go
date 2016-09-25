package util

import "net/url"

// AbsoluteURL returns the absolute URL of ref in the context of rawurl.
func AbsoluteURL(rawurl, ref string) (string, error) {
	pURL, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	pRef, err := pURL.Parse(ref)
	if err != nil {
		return "", err
	}
	return pRef.String(), nil
}
