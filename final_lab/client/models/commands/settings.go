package commands

type SettingsCommand Command

type ChangeSettingCommand SettingsCommand

const (
	ChangeSetting SettingsCommand = iota
	ViewSettings
	SaveSettings
	QuitSettings
	SettingsHelp
	InvalidSettingsCommand
)

const (
	ChangeDifficulty ChangeSettingCommand = iota
	ChangeMeasures
	QuitChangeSettings
	InvalidChangeSettingCommand
)
