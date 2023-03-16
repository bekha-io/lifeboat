package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type Provision struct {
	ProvisionType ProvisionType
	Description   string
	Callback      func()
}

func (p Provision) String() string {
	return p.Title()
}

func (p Provision) Title() string {
	return fmt.Sprintf("%v (%v)", string(p.ProvisionType), p.Description)
}

func (p Provision) CardType() CardType {
	return ProvisionCardType
}

type ProvisionType string

var (
	PaintingProvisionType    ProvisionType = "Картина"
	MoneyProvisionType       ProvisionType = "Деньги"
	JewelryProvisionType     ProvisionType = "Украшения"
	MedicalKitProvisionType  ProvisionType = "Аптечка"
	WaterProvisionType       ProvisionType = "Вода"
	LiquidWaterProvisionType ProvisionType = "Огненная вода"
	FlareGunProvisionType    ProvisionType = "Сигнальный пистолет"
	UmbrellaProvisionType    ProvisionType = "Зонтик"
	SharkBaitProvisionType   ProvisionType = "Приманка для акул"
	LifelineProvisionType    ProvisionType = "Спасательный круг"
	StickProvisionType       ProvisionType = "Дубинка"
	KnifeProvisionType       ProvisionType = "Нож"
	CompassProvisionType     ProvisionType = "Компас"
	PaddleProvisionType      ProvisionType = "Весло"
	CannibalismProvisionType ProvisionType = "Людоедство"
	HookProvisionType        ProvisionType = "Багор"
)

var MoneyProvision = Provision{
	ProvisionType: MoneyProvisionType,
	Description:   "Дает 1 очко",
	Callback: func() {

	},
}

var JewelryProvision = Provision{
	ProvisionType: JewelryProvisionType,
	Description:   "Дает 2 очка",
	Callback: func() {

	},
}

var PaintingProvision = Provision{
	ProvisionType: PaintingProvisionType,
	Description:   "Дает 3 очка",
	Callback: func() {

	},
}

var MedicalKitProvision = Provision{
	ProvisionType: MedicalKitProvisionType,
	Description:   "Можно убрать один маркер ранения с любого персонажа, включая себя. Сбросьте эту карту после использования",
	Callback: func() {

	},
}

var WaterProvision = Provision{
	ProvisionType: WaterProvisionType,
	Description:   "Используйте для того, чтобы избежать жажды во время разыгрывания карты навигации",
	Callback: func() {

	},
}

var LiquidWaterProvision = Provision{
	ProvisionType: LiquidWaterProvisionType,
	Description: "+3 к силе. Если огенная вода у вас в руках, вы можете увеличить силу на 3 до конца хода. " +
		"В конце хода вы будете испытывать дополнительную жажду. Эта карта не сбрасывается после использования",
	Callback: func() {

	},
}

var FlareGunProvision = Provision{
	ProvisionType: FlareGunProvisionType,
	Description: "+ 8 к силе. Можно однократно использовать как оружие или же выстрелить им в воздух, чтобы потянуть " +
		"3 карты навигации и сразу же получить эффект от всех увиденных птиц. Сбросьте эту карту после использования",
	Callback: func() {

	},
}

var UmbrellaProvision = Provision{
	ProvisionType: UmbrellaProvisionType,
	Description: "Откройте зонтик. Во время каждого хода, пока он у вас в руках, вы можете защищиться " +
		"от одного ранения, вызванного жаждой",
	Callback: func() {

	},
}

var SharkBaitProvision = Provision{
	ProvisionType: SharkBaitProvisionType,
	Description: "Используйте во время навигации, чтобы приманить акул. " +
		"Каждый за бортом получает по одной дополнительной ране",
	Callback: func() {

	},
}

var LifelineProvision = Provision{
	ProvisionType: LifelineProvisionType,
	Description: "Защищает от одного ранения при падении за борт. " +
		"Его можно бросить персонажу за бортом, находящемуся в сознании",
	Callback: func() {

	},
}

var StickProvision = Provision{
	ProvisionType: StickProvisionType,
	Description:   "+2 к силе",
	Callback: func() {

	},
}

var KnifeProvision = Provision{
	ProvisionType: KnifeProvisionType,
	Description:   "+3 к силе",
	Callback: func() {

	},
}

var CompassProvision = Provision{
	ProvisionType: CompassProvisionType,
	Description: "Если вы сидите на корме, то добавьте одну карту из колоды навигации на навигационный " +
		"стол перед тем, как начать выбирать на нем карту",
	Callback: func() {

	},
}

var CannibalismProvision = Provision{
	ProvisionType: CannibalismProvisionType,
	Description: "Во время совершения действий вы можете потратить ход на приготовление мертвого тела: это позволит всем на лодке вылечить одну рану. Любой несогласный может драться с вами, чтобы не допустить людоедства. " +
		"После использования сбросьте эту карту",
	Callback: func() {

	},
}

var HookProvision = Provision{
	ProvisionType: HookProvisionType,
	Description:   "+4 к силе",
	Callback: func() {

	},
}

var PaddleProvision = Provision{
	ProvisionType: PaddleProvisionType,
	Description: "+1 к силе. Если им погрести, то вы можете потянуть дополнительную карту навигации. " +
		"Можно использовать как оружие",
	Callback: func() {

	},
}

// GetProvisionAll возвращает неиспользованную колоду припасов
func GetProvisionAll() []Provision {
	deck := []Provision{
		WaterProvision, WaterProvision, WaterProvision, WaterProvision, WaterProvision, WaterProvision, WaterProvision,
		WaterProvision, WaterProvision, WaterProvision, WaterProvision, WaterProvision, WaterProvision, WaterProvision,
		WaterProvision, WaterProvision,
		LiquidWaterProvision, LiquidWaterProvision, LiquidWaterProvision,
		MedicalKitProvision, MedicalKitProvision, MedicalKitProvision, MedicalKitProvision,
		PaddleProvision, PaddleProvision,
		HookProvision,
		CannibalismProvision,
		CompassProvision,
		KnifeProvision,
		StickProvision,
		UmbrellaProvision,
		FlareGunProvision,
		MoneyProvision, MoneyProvision, MoneyProvision, MoneyProvision, MoneyProvision, MoneyProvision, MoneyProvision,
		SharkBaitProvision, SharkBaitProvision,
		PaintingProvision, PaintingProvision, PaintingProvision,
		JewelryProvision, JewelryProvision, JewelryProvision,
		LifelineProvision,
	}

	rand.Seed(time.Now().UnixMicro())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}
