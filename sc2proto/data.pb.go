// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2clientprotocol/data.proto

package SC2APIProtocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Attribute int32

const (
	Attribute_Light      Attribute = 1
	Attribute_Armored    Attribute = 2
	Attribute_Biological Attribute = 3
	Attribute_Mechanical Attribute = 4
	Attribute_Robotic    Attribute = 5
	Attribute_Psionic    Attribute = 6
	Attribute_Massive    Attribute = 7
	Attribute_Structure  Attribute = 8
	Attribute_Hover      Attribute = 9
	Attribute_Heroic     Attribute = 10
	Attribute_Summoned   Attribute = 11
)

var Attribute_name = map[int32]string{
	1:  "Light",
	2:  "Armored",
	3:  "Biological",
	4:  "Mechanical",
	5:  "Robotic",
	6:  "Psionic",
	7:  "Massive",
	8:  "Structure",
	9:  "Hover",
	10: "Heroic",
	11: "Summoned",
}
var Attribute_value = map[string]int32{
	"Light":      1,
	"Armored":    2,
	"Biological": 3,
	"Mechanical": 4,
	"Robotic":    5,
	"Psionic":    6,
	"Massive":    7,
	"Structure":  8,
	"Hover":      9,
	"Heroic":     10,
	"Summoned":   11,
}

func (x Attribute) Enum() *Attribute {
	p := new(Attribute)
	*p = x
	return p
}
func (x Attribute) String() string {
	return proto.EnumName(Attribute_name, int32(x))
}
func (x *Attribute) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Attribute_value, data, "Attribute")
	if err != nil {
		return err
	}
	*x = Attribute(value)
	return nil
}
func (Attribute) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type AbilityData_Target int32

const (
	AbilityData_None        AbilityData_Target = 1
	AbilityData_Point       AbilityData_Target = 2
	AbilityData_Unit        AbilityData_Target = 3
	AbilityData_PointOrUnit AbilityData_Target = 4
	AbilityData_PointOrNone AbilityData_Target = 5
)

var AbilityData_Target_name = map[int32]string{
	1: "None",
	2: "Point",
	3: "Unit",
	4: "PointOrUnit",
	5: "PointOrNone",
}
var AbilityData_Target_value = map[string]int32{
	"None":        1,
	"Point":       2,
	"Unit":        3,
	"PointOrUnit": 4,
	"PointOrNone": 5,
}

func (x AbilityData_Target) Enum() *AbilityData_Target {
	p := new(AbilityData_Target)
	*p = x
	return p
}
func (x AbilityData_Target) String() string {
	return proto.EnumName(AbilityData_Target_name, int32(x))
}
func (x *AbilityData_Target) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AbilityData_Target_value, data, "AbilityData_Target")
	if err != nil {
		return err
	}
	*x = AbilityData_Target(value)
	return nil
}
func (AbilityData_Target) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type Weapon_TargetType int32

const (
	Weapon_Ground Weapon_TargetType = 1
	Weapon_Air    Weapon_TargetType = 2
	Weapon_Any    Weapon_TargetType = 3
)

var Weapon_TargetType_name = map[int32]string{
	1: "Ground",
	2: "Air",
	3: "Any",
}
var Weapon_TargetType_value = map[string]int32{
	"Ground": 1,
	"Air":    2,
	"Any":    3,
}

func (x Weapon_TargetType) Enum() *Weapon_TargetType {
	p := new(Weapon_TargetType)
	*p = x
	return p
}
func (x Weapon_TargetType) String() string {
	return proto.EnumName(Weapon_TargetType_name, int32(x))
}
func (x *Weapon_TargetType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Weapon_TargetType_value, data, "Weapon_TargetType")
	if err != nil {
		return err
	}
	*x = Weapon_TargetType(value)
	return nil
}
func (Weapon_TargetType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

// May not relevant: queueable (everything is queueable).
// May not be important: AbilSetId - marine stim, marauder stim.
// Stuff omitted: transient.
// Stuff that may be important: cost, range, Alignment, targetfilters.
type AbilityData struct {
	AbilityId          *uint32             `protobuf:"varint,1,opt,name=ability_id,json=abilityId" json:"ability_id,omitempty"`
	LinkName           *string             `protobuf:"bytes,2,opt,name=link_name,json=linkName" json:"link_name,omitempty"`
	LinkIndex          *uint32             `protobuf:"varint,3,opt,name=link_index,json=linkIndex" json:"link_index,omitempty"`
	ButtonName         *string             `protobuf:"bytes,4,opt,name=button_name,json=buttonName" json:"button_name,omitempty"`
	FriendlyName       *string             `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName" json:"friendly_name,omitempty"`
	Hotkey             *string             `protobuf:"bytes,6,opt,name=hotkey" json:"hotkey,omitempty"`
	RemapsToAbilityId  *uint32             `protobuf:"varint,7,opt,name=remaps_to_ability_id,json=remapsToAbilityId" json:"remaps_to_ability_id,omitempty"`
	Available          *bool               `protobuf:"varint,8,opt,name=available" json:"available,omitempty"`
	Target             *AbilityData_Target `protobuf:"varint,9,opt,name=target,enum=SC2APIProtocol.AbilityData_Target" json:"target,omitempty"`
	AllowMinimap       *bool               `protobuf:"varint,10,opt,name=allow_minimap,json=allowMinimap" json:"allow_minimap,omitempty"`
	AllowAutocast      *bool               `protobuf:"varint,11,opt,name=allow_autocast,json=allowAutocast" json:"allow_autocast,omitempty"`
	IsBuilding         *bool               `protobuf:"varint,12,opt,name=is_building,json=isBuilding" json:"is_building,omitempty"`
	FootprintRadius    *float32            `protobuf:"fixed32,13,opt,name=footprint_radius,json=footprintRadius" json:"footprint_radius,omitempty"`
	IsInstantPlacement *bool               `protobuf:"varint,14,opt,name=is_instant_placement,json=isInstantPlacement" json:"is_instant_placement,omitempty"`
	CastRange          *float32            `protobuf:"fixed32,15,opt,name=cast_range,json=castRange" json:"cast_range,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *AbilityData) Reset()                    { *m = AbilityData{} }
func (m *AbilityData) String() string            { return proto.CompactTextString(m) }
func (*AbilityData) ProtoMessage()               {}
func (*AbilityData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *AbilityData) GetAbilityId() uint32 {
	if m != nil && m.AbilityId != nil {
		return *m.AbilityId
	}
	return 0
}

func (m *AbilityData) GetLinkName() string {
	if m != nil && m.LinkName != nil {
		return *m.LinkName
	}
	return ""
}

func (m *AbilityData) GetLinkIndex() uint32 {
	if m != nil && m.LinkIndex != nil {
		return *m.LinkIndex
	}
	return 0
}

func (m *AbilityData) GetButtonName() string {
	if m != nil && m.ButtonName != nil {
		return *m.ButtonName
	}
	return ""
}

func (m *AbilityData) GetFriendlyName() string {
	if m != nil && m.FriendlyName != nil {
		return *m.FriendlyName
	}
	return ""
}

func (m *AbilityData) GetHotkey() string {
	if m != nil && m.Hotkey != nil {
		return *m.Hotkey
	}
	return ""
}

func (m *AbilityData) GetRemapsToAbilityId() uint32 {
	if m != nil && m.RemapsToAbilityId != nil {
		return *m.RemapsToAbilityId
	}
	return 0
}

func (m *AbilityData) GetAvailable() bool {
	if m != nil && m.Available != nil {
		return *m.Available
	}
	return false
}

func (m *AbilityData) GetTarget() AbilityData_Target {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return AbilityData_None
}

func (m *AbilityData) GetAllowMinimap() bool {
	if m != nil && m.AllowMinimap != nil {
		return *m.AllowMinimap
	}
	return false
}

func (m *AbilityData) GetAllowAutocast() bool {
	if m != nil && m.AllowAutocast != nil {
		return *m.AllowAutocast
	}
	return false
}

func (m *AbilityData) GetIsBuilding() bool {
	if m != nil && m.IsBuilding != nil {
		return *m.IsBuilding
	}
	return false
}

func (m *AbilityData) GetFootprintRadius() float32 {
	if m != nil && m.FootprintRadius != nil {
		return *m.FootprintRadius
	}
	return 0
}

func (m *AbilityData) GetIsInstantPlacement() bool {
	if m != nil && m.IsInstantPlacement != nil {
		return *m.IsInstantPlacement
	}
	return false
}

func (m *AbilityData) GetCastRange() float32 {
	if m != nil && m.CastRange != nil {
		return *m.CastRange
	}
	return 0
}

type DamageBonus struct {
	Attribute        *Attribute `protobuf:"varint,1,opt,name=attribute,enum=SC2APIProtocol.Attribute" json:"attribute,omitempty"`
	Bonus            *float32   `protobuf:"fixed32,2,opt,name=bonus" json:"bonus,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *DamageBonus) Reset()                    { *m = DamageBonus{} }
func (m *DamageBonus) String() string            { return proto.CompactTextString(m) }
func (*DamageBonus) ProtoMessage()               {}
func (*DamageBonus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *DamageBonus) GetAttribute() Attribute {
	if m != nil && m.Attribute != nil {
		return *m.Attribute
	}
	return Attribute_Light
}

func (m *DamageBonus) GetBonus() float32 {
	if m != nil && m.Bonus != nil {
		return *m.Bonus
	}
	return 0
}

type Weapon struct {
	Type             *Weapon_TargetType `protobuf:"varint,1,opt,name=type,enum=SC2APIProtocol.Weapon_TargetType" json:"type,omitempty"`
	Damage           *float32           `protobuf:"fixed32,2,opt,name=damage" json:"damage,omitempty"`
	DamageBonus      []*DamageBonus     `protobuf:"bytes,3,rep,name=damage_bonus,json=damageBonus" json:"damage_bonus,omitempty"`
	Attacks          *uint32            `protobuf:"varint,4,opt,name=attacks" json:"attacks,omitempty"`
	Range            *float32           `protobuf:"fixed32,5,opt,name=range" json:"range,omitempty"`
	Speed            *float32           `protobuf:"fixed32,6,opt,name=speed" json:"speed,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Weapon) Reset()                    { *m = Weapon{} }
func (m *Weapon) String() string            { return proto.CompactTextString(m) }
func (*Weapon) ProtoMessage()               {}
func (*Weapon) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Weapon) GetType() Weapon_TargetType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Weapon_Ground
}

func (m *Weapon) GetDamage() float32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func (m *Weapon) GetDamageBonus() []*DamageBonus {
	if m != nil {
		return m.DamageBonus
	}
	return nil
}

func (m *Weapon) GetAttacks() uint32 {
	if m != nil && m.Attacks != nil {
		return *m.Attacks
	}
	return 0
}

func (m *Weapon) GetRange() float32 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

func (m *Weapon) GetSpeed() float32 {
	if m != nil && m.Speed != nil {
		return *m.Speed
	}
	return 0
}

type UnitTypeData struct {
	UnitId      *uint32 `protobuf:"varint,1,opt,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Available   *bool   `protobuf:"varint,3,opt,name=available" json:"available,omitempty"`
	CargoSize   *uint32 `protobuf:"varint,4,opt,name=cargo_size,json=cargoSize" json:"cargo_size,omitempty"`
	MineralCost *uint32 `protobuf:"varint,12,opt,name=mineral_cost,json=mineralCost" json:"mineral_cost,omitempty"`
	VespeneCost *uint32 `protobuf:"varint,13,opt,name=vespene_cost,json=vespeneCost" json:"vespene_cost,omitempty"`
	// Values include changes from upgrades
	Attributes       []Attribute `protobuf:"varint,8,rep,name=attributes,enum=SC2APIProtocol.Attribute" json:"attributes,omitempty"`
	MovementSpeed    *float32    `protobuf:"fixed32,9,opt,name=movement_speed,json=movementSpeed" json:"movement_speed,omitempty"`
	Armor            *float32    `protobuf:"fixed32,10,opt,name=armor" json:"armor,omitempty"`
	Weapons          []*Weapon   `protobuf:"bytes,11,rep,name=weapons" json:"weapons,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *UnitTypeData) Reset()                    { *m = UnitTypeData{} }
func (m *UnitTypeData) String() string            { return proto.CompactTextString(m) }
func (*UnitTypeData) ProtoMessage()               {}
func (*UnitTypeData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *UnitTypeData) GetUnitId() uint32 {
	if m != nil && m.UnitId != nil {
		return *m.UnitId
	}
	return 0
}

func (m *UnitTypeData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UnitTypeData) GetAvailable() bool {
	if m != nil && m.Available != nil {
		return *m.Available
	}
	return false
}

func (m *UnitTypeData) GetCargoSize() uint32 {
	if m != nil && m.CargoSize != nil {
		return *m.CargoSize
	}
	return 0
}

func (m *UnitTypeData) GetMineralCost() uint32 {
	if m != nil && m.MineralCost != nil {
		return *m.MineralCost
	}
	return 0
}

func (m *UnitTypeData) GetVespeneCost() uint32 {
	if m != nil && m.VespeneCost != nil {
		return *m.VespeneCost
	}
	return 0
}

func (m *UnitTypeData) GetAttributes() []Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *UnitTypeData) GetMovementSpeed() float32 {
	if m != nil && m.MovementSpeed != nil {
		return *m.MovementSpeed
	}
	return 0
}

func (m *UnitTypeData) GetArmor() float32 {
	if m != nil && m.Armor != nil {
		return *m.Armor
	}
	return 0
}

func (m *UnitTypeData) GetWeapons() []*Weapon {
	if m != nil {
		return m.Weapons
	}
	return nil
}

type UpgradeData struct {
	UpgradeId        *uint32 `protobuf:"varint,1,opt,name=upgrade_id,json=upgradeId" json:"upgrade_id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpgradeData) Reset()                    { *m = UpgradeData{} }
func (m *UpgradeData) String() string            { return proto.CompactTextString(m) }
func (*UpgradeData) ProtoMessage()               {}
func (*UpgradeData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UpgradeData) GetUpgradeId() uint32 {
	if m != nil && m.UpgradeId != nil {
		return *m.UpgradeId
	}
	return 0
}

func (m *UpgradeData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type BuffData struct {
	BuffId           *uint32 `protobuf:"varint,1,opt,name=buff_id,json=buffId" json:"buff_id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BuffData) Reset()                    { *m = BuffData{} }
func (m *BuffData) String() string            { return proto.CompactTextString(m) }
func (*BuffData) ProtoMessage()               {}
func (*BuffData) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *BuffData) GetBuffId() uint32 {
	if m != nil && m.BuffId != nil {
		return *m.BuffId
	}
	return 0
}

func (m *BuffData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AbilityData)(nil), "SC2APIProtocol.AbilityData")
	proto.RegisterType((*DamageBonus)(nil), "SC2APIProtocol.DamageBonus")
	proto.RegisterType((*Weapon)(nil), "SC2APIProtocol.Weapon")
	proto.RegisterType((*UnitTypeData)(nil), "SC2APIProtocol.UnitTypeData")
	proto.RegisterType((*UpgradeData)(nil), "SC2APIProtocol.UpgradeData")
	proto.RegisterType((*BuffData)(nil), "SC2APIProtocol.BuffData")
	proto.RegisterEnum("SC2APIProtocol.Attribute", Attribute_name, Attribute_value)
	proto.RegisterEnum("SC2APIProtocol.AbilityData_Target", AbilityData_Target_name, AbilityData_Target_value)
	proto.RegisterEnum("SC2APIProtocol.Weapon_TargetType", Weapon_TargetType_name, Weapon_TargetType_value)
}

func init() { proto.RegisterFile("s2clientprotocol/data.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 890 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x21, 0xf9, 0x53, 0x47, 0xb6, 0xa3, 0x11, 0x41, 0xa7, 0x21, 0x1b, 0xe6, 0x7a, 0x18,
	0xe0, 0xf5, 0x22, 0x2d, 0x02, 0x0c, 0xc5, 0x76, 0x31, 0xcc, 0x69, 0x81, 0xd5, 0xc0, 0xd2, 0x19,
	0x72, 0x8a, 0xdd, 0x0c, 0x10, 0x68, 0x89, 0x76, 0x88, 0x48, 0xa4, 0x40, 0x52, 0xee, 0xdc, 0xfb,
	0x5d, 0xec, 0x25, 0xf6, 0x7c, 0x7b, 0x8c, 0x81, 0x87, 0x72, 0xa2, 0x66, 0x41, 0xef, 0xf4, 0xff,
	0xf1, 0x7f, 0x68, 0x9e, 0x2f, 0xc3, 0x99, 0xbe, 0xc8, 0x0a, 0xce, 0x84, 0xa9, 0x94, 0x34, 0x32,
	0x93, 0xc5, 0xf3, 0x9c, 0x1a, 0x7a, 0x8e, 0x8a, 0x4c, 0xd6, 0xaf, 0x2e, 0x16, 0xab, 0xe5, 0xaa,
	0x39, 0x9a, 0xfd, 0xd5, 0x83, 0x70, 0xb1, 0xe1, 0x05, 0x37, 0x87, 0xd7, 0xd4, 0x50, 0xf2, 0x15,
	0x00, 0x75, 0x32, 0xe5, 0x79, 0xec, 0x4d, 0xbd, 0xf9, 0x38, 0x09, 0x1a, 0xb2, 0xcc, 0xc9, 0x19,
	0x04, 0x05, 0x17, 0xb7, 0xa9, 0xa0, 0x25, 0x8b, 0xfd, 0xa9, 0x37, 0x0f, 0x92, 0xa1, 0x05, 0x6f,
	0x69, 0xc9, 0x6c, 0x2c, 0x1e, 0x72, 0x91, 0xb3, 0x3f, 0xe3, 0x8e, 0x8b, 0xb5, 0x64, 0x69, 0x01,
	0xf9, 0x1a, 0xc2, 0x4d, 0x6d, 0x8c, 0x14, 0x2e, 0xba, 0x8b, 0xd1, 0xe0, 0x10, 0xc6, 0x7f, 0x03,
	0xe3, 0xad, 0xe2, 0x4c, 0xe4, 0xc5, 0xc1, 0x59, 0x7a, 0x68, 0x19, 0x1d, 0x21, 0x9a, 0x9e, 0x40,
	0xff, 0x46, 0x9a, 0x5b, 0x76, 0x88, 0xfb, 0x78, 0xda, 0x28, 0xf2, 0x1c, 0x4e, 0x15, 0x2b, 0x69,
	0xa5, 0x53, 0x23, 0xd3, 0x56, 0x0a, 0x03, 0x7c, 0xc6, 0x67, 0xee, 0xec, 0x5a, 0x2e, 0xee, 0x52,
	0xf9, 0x12, 0x02, 0xba, 0xa7, 0xbc, 0xa0, 0x9b, 0x82, 0xc5, 0xc3, 0xa9, 0x37, 0x1f, 0x26, 0xf7,
	0x80, 0xfc, 0x08, 0x7d, 0x43, 0xd5, 0x8e, 0x99, 0x38, 0x98, 0x7a, 0xf3, 0xc9, 0xc5, 0xec, 0xfc,
	0xe3, 0xc2, 0x9d, 0xb7, 0x8a, 0x76, 0x7e, 0x8d, 0xce, 0xa4, 0x89, 0xb0, 0x79, 0xd0, 0xa2, 0x90,
	0xef, 0xd3, 0x92, 0x0b, 0x5e, 0xd2, 0x2a, 0x06, 0xbc, 0x7d, 0x84, 0xf0, 0xca, 0x31, 0xf2, 0x2d,
	0x4c, 0x9c, 0x89, 0xd6, 0x46, 0x66, 0x54, 0x9b, 0x38, 0x44, 0x97, 0x0b, 0x5d, 0x34, 0xd0, 0x16,
	0x8d, 0xeb, 0x74, 0x53, 0xf3, 0x22, 0xe7, 0x62, 0x17, 0x8f, 0xd0, 0x03, 0x5c, 0x5f, 0x36, 0x84,
	0x7c, 0x07, 0xd1, 0x56, 0x4a, 0x53, 0x29, 0x2e, 0x4c, 0xaa, 0x68, 0xce, 0x6b, 0x1d, 0x8f, 0xa7,
	0xde, 0xdc, 0x4f, 0x4e, 0xee, 0x78, 0x82, 0x98, 0xbc, 0x80, 0x53, 0xae, 0x53, 0x2e, 0xb4, 0xa1,
	0xc2, 0xa4, 0x55, 0x41, 0x33, 0x56, 0x32, 0x61, 0xe2, 0x09, 0x5e, 0x4a, 0xb8, 0x5e, 0xba, 0xa3,
	0xd5, 0xf1, 0xc4, 0x76, 0xd4, 0xbe, 0x22, 0x55, 0x54, 0xec, 0x58, 0x7c, 0x82, 0xd7, 0x06, 0x96,
	0x24, 0x16, 0xcc, 0x96, 0xd0, 0x77, 0xa9, 0x93, 0x21, 0x74, 0xdf, 0x4a, 0xc1, 0x22, 0x8f, 0x04,
	0xd0, 0x5b, 0x49, 0x2e, 0x4c, 0xe4, 0x5b, 0xf8, 0x4e, 0x70, 0x13, 0x75, 0xc8, 0x09, 0x84, 0x08,
	0x7f, 0x53, 0x08, 0xba, 0x2d, 0x80, 0x61, 0xbd, 0xd9, 0x1f, 0x10, 0xbe, 0xa6, 0x25, 0xdd, 0xb1,
	0x4b, 0x29, 0x6a, 0x4d, 0x5e, 0x42, 0x40, 0x8d, 0x51, 0x7c, 0x53, 0x1b, 0x86, 0x53, 0x38, 0xb9,
	0xf8, 0xe2, 0x7f, 0x1d, 0x38, 0x1a, 0x92, 0x7b, 0x2f, 0x39, 0x85, 0xde, 0xc6, 0xde, 0x80, 0xc3,
	0xe9, 0x27, 0x4e, 0xcc, 0xfe, 0xf6, 0xa1, 0xff, 0x3b, 0xa3, 0x95, 0x14, 0xe4, 0x7b, 0xe8, 0x9a,
	0x43, 0x75, 0xbc, 0xf4, 0xe9, 0xc3, 0x4b, 0x9d, 0xab, 0xe9, 0xe8, 0xf5, 0xa1, 0x62, 0x09, 0xda,
	0xed, 0xd8, 0xe5, 0xf8, 0xbe, 0xe6, 0xe2, 0x46, 0x91, 0x9f, 0x60, 0xe4, 0xbe, 0x52, 0xf7, 0xb3,
	0x9d, 0x69, 0x67, 0x1e, 0x5e, 0x9c, 0x3d, 0xbc, 0xb6, 0x95, 0x5b, 0x12, 0xe6, 0xad, 0x44, 0x63,
	0x18, 0x50, 0x63, 0x68, 0x76, 0xab, 0x71, 0x21, 0xc6, 0xc9, 0x51, 0xda, 0x4c, 0x5c, 0xd9, 0x7b,
	0x2e, 0x13, 0x14, 0x96, 0xea, 0x8a, 0xb1, 0x1c, 0xa7, 0xdf, 0x4f, 0x9c, 0x98, 0x3d, 0x03, 0xb8,
	0x7f, 0x31, 0x01, 0xe8, 0xff, 0xa2, 0x64, 0x2d, 0xf2, 0xc8, 0x23, 0x03, 0xe8, 0x2c, 0xb8, 0x8a,
	0x7c, 0xfc, 0x10, 0x87, 0xa8, 0x33, 0xfb, 0xd7, 0x87, 0x91, 0xed, 0x82, 0xb5, 0xe2, 0xca, 0x7f,
	0x0e, 0x83, 0x5a, 0x70, 0x73, 0xbf, 0xef, 0x7d, 0x2b, 0x97, 0x39, 0x21, 0xd0, 0x6d, 0xed, 0x39,
	0x7e, 0x7f, 0xbc, 0x35, 0x9d, 0x87, 0x5b, 0x83, 0xf3, 0xa2, 0x76, 0x32, 0xd5, 0xfc, 0x03, 0x6b,
	0x12, 0x0a, 0x90, 0xac, 0xf9, 0x07, 0x46, 0x9e, 0xc2, 0xa8, 0xe4, 0x82, 0x29, 0x5a, 0xa4, 0x99,
	0xd4, 0x06, 0xa7, 0x79, 0x9c, 0x84, 0x0d, 0x7b, 0x25, 0xb5, 0xb1, 0x96, 0x3d, 0xd3, 0x15, 0x13,
	0xcc, 0x59, 0xc6, 0xce, 0xd2, 0x30, 0xb4, 0xfc, 0x00, 0x70, 0xd7, 0x6f, 0x1d, 0x0f, 0xa7, 0x9d,
	0x4f, 0x0f, 0x47, 0xcb, 0x6c, 0x97, 0xae, 0x94, 0x7b, 0x9c, 0xed, 0xd4, 0x95, 0x31, 0xc0, 0x32,
	0x8e, 0x8f, 0x74, 0x6d, 0xa1, 0x2d, 0x32, 0x55, 0xa5, 0x54, 0xb8, 0xb8, 0x7e, 0xe2, 0x04, 0x79,
	0x01, 0x83, 0xf7, 0x38, 0x1d, 0x3a, 0x0e, 0xb1, 0xcb, 0x4f, 0x1e, 0x1f, 0x9e, 0xe4, 0x68, 0x9b,
	0xfd, 0x0c, 0xe1, 0xbb, 0x6a, 0xa7, 0x68, 0xce, 0x8e, 0xff, 0xad, 0xb5, 0x93, 0xad, 0xff, 0xd6,
	0x86, 0x3c, 0x5e, 0xee, 0xd9, 0x4b, 0x18, 0x5e, 0xd6, 0xdb, 0xed, 0xb1, 0x4f, 0x9b, 0x7a, 0xbb,
	0x6d, 0xf5, 0xc9, 0xca, 0xc7, 0x03, 0x9f, 0xfd, 0xe3, 0x41, 0x70, 0x57, 0x03, 0xbb, 0x94, 0xbf,
	0xf2, 0xdd, 0x8d, 0x89, 0x3c, 0x12, 0xc2, 0x60, 0x61, 0xd3, 0x61, 0x79, 0xe4, 0x93, 0x09, 0xc0,
	0x25, 0x97, 0x85, 0xdc, 0xf1, 0x8c, 0x16, 0x51, 0xc7, 0xea, 0x2b, 0x96, 0xdd, 0x50, 0x81, 0xba,
	0x6b, 0xcd, 0x89, 0xdc, 0x48, 0xc3, 0xb3, 0xa8, 0x67, 0xc5, 0x4a, 0x73, 0x29, 0x78, 0x16, 0xf5,
	0xad, 0xb8, 0xa2, 0x5a, 0xf3, 0x3d, 0x8b, 0x06, 0x64, 0x0c, 0xc1, 0xda, 0xa8, 0x3a, 0x33, 0xb5,
	0x62, 0xd1, 0xd0, 0xfe, 0xda, 0x1b, 0xb9, 0x67, 0x2a, 0x0a, 0xec, 0x28, 0xbe, 0x61, 0x4a, 0xf2,
	0x2c, 0x02, 0x32, 0x82, 0xe1, 0xba, 0x2e, 0x4b, 0x29, 0x58, 0x1e, 0x85, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x68, 0x31, 0x1b, 0x77, 0xa6, 0x06, 0x00, 0x00,
}