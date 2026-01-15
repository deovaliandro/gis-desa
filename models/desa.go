package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Desa struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NAMOBJ            string             `bson:"NAMOBJ"`
	FCODE             string             `bson:"FCODE"`
	REMARK            string             `bson:"REMARK"`
	METADATA          string             `bson:"METADATA"`
	SRS_ID            string             `bson:"SRS_ID"`
	KDCPUM            string             `bson:"KDCPUM"`
	KDEPUM            string             `bson:"KDEPUM"`
	KDPKAB            string             `bson:"KDPKAB"`
	KDPPUM            string             `bson:"KDPPUM"`
	LUASWH            float64            `bson:"LUASWH"`
	TIPADM            int                `bson:"TIPADM"`
	WADMKC            string             `bson:"WADMKC"`
	WADMKD            string             `bson:"WADMKD"`
	WADMKK            string             `bson:"WADMKK"`
	WADMPR            string             `bson:"WADMPR"`
	UUPP              string             `bson:"UUPP"`
	ShapeLeng         float64            `bson:"Shape_Leng"`
	ShapeArea         float64            `bson:"Shape_Area"`
	TingkatPendidikan int                `bson:"TINGKAT_PENDIDIKAN"`
	Kesehatan         int                `bson:"KESEHATAN"`
}
