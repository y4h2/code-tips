package example

type Cloneable interface{}

type WeeklyLog struct {
	Cloneable

	name    string
	date    string
	content string
}

func (log *WeeklyLog) SetDate(date string) {
	log.date = date
}

func (log *WeeklyLog) GetDate() string {
	return log.date
}

func (log *WeeklyLog) SetName(name string) {
	log.name = name
}

func (log *WeeklyLog) GetName() string {
	return log.name
}

func (log *WeeklyLog) SetContent(content string) {
	log.content = content
}

func (log *WeeklyLog) GetContent() string {
	return log.content
}

func (log *WeeklyLog) Clone() *WeeklyLog {
	return &WeeklyLog{
		name:    log.name,
		date:    log.date,
		content: log.content,
	}
}
