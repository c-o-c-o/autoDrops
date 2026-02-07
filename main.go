package main

import (
	"autoDrops/data"
	"autoDrops/env"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	edr, err := env.GetExecDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	app := &cli.App{
		Name:            "autoDrops",
		Usage:           Version,
		Description:     "設定に従ってaviutlにファイルをドロップします",
		Version:         Version,
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "text",
				Aliases:  []string{"t"},
				Required: true,
				Usage:    "パターンマッチするテキストファイル",
			},
			&cli.PathFlag{
				Name:    "profile",
				Aliases: []string{"p"},
				Value:   filepath.Join(edr, "profile.yml"),
				Usage:   "プロファイルのパス",
			},
		},
		Action: appfunc,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func appfunc(c *cli.Context) error {
	bprof, err := os.ReadFile(c.Path("profile"))
	if err != nil {
		return err
	}

	profile := &data.Profile{}
	err = yaml.Unmarshal(bprof, profile)
	if err != nil {
		return err
	}

	client := profile.Client
	if client == "" {
		client = "cdrops.exe"
	}

	cmds, err := getCommand(profile, c.Path("text"), c.Args().Slice())
	if err != nil {
		return err
	}

	edr, err := env.GetExecDir()
	if err != nil {
		return err
	}

	inst := exec.Command(filepath.Join(edr, client), cmds...)
	inst.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	err = inst.Run()
	if err != nil {
		return err
	}

	return nil
}

func getCommand(profile *data.Profile, tpath string, paths []string) ([]string, error) {
	for _, act := range profile.Actors {
		if act.Target == "" {
			act.Target = "Name"
		}

		text := ""
		switch act.Target {
		case "Name":
			text = filepath.Base(tpath)
		case "Text":
			tbyte, err := os.ReadFile(tpath)
			if err != nil {
				return nil, err
			}
			text = string(tbyte)
		default:
			return nil, errors.New("Targetが間違っています\nNameかTextを使用してください")
		}

		ptn, err := regexp.Compile(act.Pattern)
		if err != nil {
			return nil, errors.New("Patternが間違っています\n正しい正規表現を使用してください")
		}

		if ptn.MatchString(text) {
			return []string{
				strconv.Itoa(act.Layer) + "*1000*" + strings.Join(paths, "*"),
				"1*" + strconv.Itoa(profile.Advance),
			}, nil
		}
	}
	return nil, errors.New("全てのPatternにマッチしませんでした")
}
