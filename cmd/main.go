// Package main provides ...
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

var (
	homepath        = os.Getenv("HOME")
	langpluginsfile string
)

func main() {
	userplugins := homepath + "/.thinkvim.d/plugins.yaml"
	colorscheme(userplugins)
	userlangs := clilangqs()
	var wg sync.WaitGroup
	for _, v := range userlangs {
		langpluginsfile = homepath + "/.config/nvim/lang/" + v + ".yaml"
		wg.Add(1)
		go writeuserplugins(&wg, langpluginsfile, userplugins)
	}
	wg.Wait()
	fmt.Println("generat your languages plugins.yaml success")
}

func writeuserplugins(wg *sync.WaitGroup, filename, userplugins string) {
	mu := sync.RWMutex{}
	mu.RLock()
	f, err := os.OpenFile(userplugins, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Open user plugins.yaml failed!")
	}
	content, _ := ioutil.ReadFile(filename)
	_, err = f.Write(content)
	if err != nil {
		log.Fatalln("write to user plugins.yaml failed")
		os.Exit(1)
	}
	wg.Done()
	f.Close()
	mu.RUnlock()
}

func answertemplate(questionname, message string, options []string, pagesize int) ([]string, error) {
	answers := []string{}
	var question = []*survey.Question{
		{
			Name: questionname,
			Prompt: &survey.MultiSelect{
				Message: message,
				Options: options,
			},
		},
	}
	err := survey.Ask(question, &answers, survey.WithIcons(func(icons *survey.IconSet) {
		icons.UnmarkedOption.Text = "○"
		icons.MarkedOption.Text = "◉"
	}), survey.WithPageSize(pagesize))
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func colorscheme(filename string) {
	questionname := "Colorscheme Question"
	message := "Choose your favorite colorscheme"
	pagesize := 19
	options := []string{
		"hardcoreplayers/oceanic-material(support by default)",
		"drewtempelmeyer/palenight.vim",
		"ayu-theme/ayu-vim",
		"NLKNguyen/papercolor-theme",
		"lifepillar/vim-gruvbox8",
		"lifepillar/vim-solarized8",
		"joshdick/onedark.vim",
		"arcticicestudio/nord-vim",
		"rakr/vim-one",
		"mhartington/oceanic-next",
		"patstockwell/vim-monokai-tasty",
		"dracula/vim",
		"chriskempson/base16-vim",
		"kristijanhusak/vim-hybrid-material",
		"kyoz/purify",
		"nanotech/jellybeans.vim",
	}
	colorschemes, err := answertemplate(questionname, message, options, pagesize)
	if err != nil {
		log.Fatalf("Colorscehme Question error:%s", err)
		os.Exit(1)
	}
	var wg sync.WaitGroup
	for _, v := range colorschemes {
		if v != "hardcoreplayers/oceanic-material(support by default)" {
			wg.Add(1)
			repo := fmt.Sprintf("- { repo: %s }\n", v)
			go func(wg *sync.WaitGroup) {
				var mu sync.RWMutex
				mu.RLock()
				defer mu.RUnlock()
				f, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				defer f.Close()
				_, err := f.WriteString(repo)
				if err != nil {
					log.Fatalf("Write colorscheme to %s failed error:%s", filename, err)
					os.Exit(1)
				}
				wg.Done()
			}(&wg)
		}
	}
	wg.Wait()
	themecache := homepath + "/.cache/vim/theme.txt"
	f, _ := os.OpenFile(themecache, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	colors := GetColorScheme(colorschemes)
	_, err = f.WriteString(colors)
	if err != nil {
		log.Fatalf("Write to theme cache file error %s", err)
		os.Exit(1)
	}
	f.Close()
	fmt.Println("generate colorscehme success")
}

func GetColorScheme(colorschemes []string) string {
	colors := ""
	switch colorschemes[0] {
	case "hardcoreplayers/oceanic-material(support by default)":
		colors = "oceanic_material\n"
	case "NLKNguyen/papercolor-theme":
		colors = "PaperColor\n"
	case "drewtempelmeyer/palenight.vim":
		colors = "palenight\n"
	case "ayu-theme/ayu-vim":
		colors = "ayu\n"
	case "lifepillar/vim-gruvbox8":
		colors = "gruvbox8\n"
	case "lifepillar/vim-solarized8":
		colors = "solarized8\n"
	case "joshdick/onedark.vim":
		colors = "onedark\n"
	case "arcticicestudio/nord-vim":
		colors = "nord\n"
	case "rakr/vim-one":
		colors = "one\n"
	case "mhartington/oceanic-next":
		colors = "OceanicNext\n"
	case "patstockwell/vim-monokai-tasty":
		colors = "vim-monokai-tasty\n"
	case "dracula/vim":
		colors = "dracula\n"
	case "chriskempson/base16-vim":
		colors = "base16-default-dark\n"
	case "kristijanhusak/vim-hybrid-material":
		colors = "hybrid_material\n"
	case "kyoz/purify":
		colors = "purify\n"
	case "nanotech/jellybeans.vim":
		colors = "jellybeans\n"
	default:
		colors = ""
	}
	return colors
}

func clilangqs() []string {
	questionname := "LanguageQuestion"
	message := "What Languages do you write"
	pagesize := 19
	options := []string{
		"c-family",
		"R",
		"javascript",
		"typescript",
		"react",
		"vue",
		"go",
		"rust",
		"haskell",
		"php",
		"ruby",
		"scala",
		"shell",
		"lua",
		"python",
		"dockerfile",
		"json",
		"nginx",
		"toml",
		"html",
		"css",
		"less",
		"sass",
		"stylus",
		"sql",
		"dart",
	}
	answers, err := answertemplate(questionname, message, options, pagesize)
	if err != nil {
		panic(err)
	}
	return answers
}
