package builder

type Actor struct {
	typ       string
	sex       string
	face      string
	costume   string
	hairstyle string
}

func NewActor() Actor {
	return Actor{}
}

func (a *Actor) SetType(typ string) {
	a.typ = typ
}

func (a *Actor) SetSex(sex string) {
	a.sex = sex
}
func (a *Actor) SetFace(face string) {
	a.face = face
}
func (a *Actor) SetCostume(costume string) {
	a.costume = costume
}
func (a *Actor) SetHairstyle(hairstyle string) {
	a.hairstyle = hairstyle
}
func (a *Actor) GetType() string {
	return a.typ
}

func (a *Actor) GetSex() string {
	return a.sex
}
func (a *Actor) GetFace() string {
	return a.face
}
func (a *Actor) GetCostume() string {
	return a.costume
}
func (a *Actor) GetHairstyle() string {
	return a.hairstyle
}

type ActorBuilder interface {
	BuildType()
	BuildSex()
	BuildFace()
	BuildCostume()
	BuildHairstyle()
	CreateActor() Actor
}

type ActorCreator struct {
	actor Actor
}

func NewActorCreator() ActorCreator {
	return ActorCreator{actor: NewActor()}
}

func (ac ActorCreator) CreateActor() Actor {
	return ac.actor
}

type HeroBuilder struct {
	ActorCreator
}

func NewHeroBuilder() *HeroBuilder {
	return &HeroBuilder{NewActorCreator()}
}

func (builder *HeroBuilder) BuildType() {
	builder.actor.SetType("hero")
}
func (builder *HeroBuilder) BuildSex() {
	builder.actor.SetSex("male")
}
func (builder *HeroBuilder) BuildFace() {
	builder.actor.SetFace("handsome")
}
func (builder *HeroBuilder) BuildCostume() {
	builder.actor.SetCostume("armor")
}
func (builder *HeroBuilder) BuildHairstyle() {
	builder.actor.SetHairstyle("curve")
}

type AngelBuilder struct {
	ActorCreator
}

func NewAngelBuilder() *AngelBuilder {
	return &AngelBuilder{NewActorCreator()}
}

func (builder *AngelBuilder) BuildType() {
	builder.actor.SetType("angel")
}
func (builder *AngelBuilder) BuildSex() {
	builder.actor.SetSex("female")
}
func (builder *AngelBuilder) BuildFace() {
	builder.actor.SetFace("beautiful")
}
func (builder *AngelBuilder) BuildCostume() {
	builder.actor.SetCostume("shorts")
}
func (builder *AngelBuilder) BuildHairstyle() {
	builder.actor.SetHairstyle("golden")
}

type DevilBuilder struct {
	ActorCreator
}

func NewDevilBuilder() *DevilBuilder {
	return &DevilBuilder{NewActorCreator()}
}

func (builder *DevilBuilder) BuildType() {
	builder.actor.SetType("devil")
}
func (builder *DevilBuilder) BuildSex() {
	builder.actor.SetSex("male")
}
func (builder *DevilBuilder) BuildFace() {
	builder.actor.SetFace("ugly")
}
func (builder *DevilBuilder) BuildCostume() {
	builder.actor.SetCostume("coat")
}
func (builder *DevilBuilder) BuildHairstyle() {
	builder.actor.SetHairstyle("bald")
}

func ConstructActor(ab ActorBuilder) Actor {
	var actor Actor
	ab.BuildType()
	ab.BuildSex()
	ab.BuildFace()
	ab.BuildCostume()
	ab.BuildHairstyle()
	actor = ab.CreateActor()
	return actor
}
