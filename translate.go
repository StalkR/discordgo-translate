package translate

import (
  "encoding/json"
  "errors"
  "fmt"
  "html"
  "net/http"
  "net/url"
)

const googleTranslateAPI = "https://www.googleapis.com/language/translate/v2"

// translate translates text using Google Translate API.
// It returns the detected source language and the translation.
func translate(apiKey, target, text string) (string, string, error) {
  params := url.Values{}
  params.Set("key", apiKey)
  // no "source" parameter to automatically detect language
  params.Set("target", target)
  params.Set("q", text)
  resp, err := http.Get(fmt.Sprintf("%s?%s", googleTranslateAPI, params.Encode()))
  if err != nil {
    return "", "", err
  }
  defer resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    var v struct {
      Error struct {
        Code    int
        Message string
      }
    }
    if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
      return "", "", err
    }
    return "", "", fmt.Errorf("%v: %v", v.Error.Code, v.Error.Message)
  }
  var v struct {
    Data struct {
      Translations []struct {
        TranslatedText         string
        DetectedSourceLanguage string
      }
    }
  }
  if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
    return "", "", err
  }
  if len(v.Data.Translations) == 0 {
    return "", "", errors.New("no translation")
  }
  source := v.Data.Translations[0].DetectedSourceLanguage
  translated := html.UnescapeString(v.Data.Translations[0].TranslatedText)
  return source, translated, nil
}
