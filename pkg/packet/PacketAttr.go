package packet

type MsgAttr uint16

const (
	MsgAttrPrivate               MsgAttr   = 1 << 0
	MsgAttrCrash                 MsgAttr   = 1 << 1
//	MsgAttrRecd                  MsgAttr   = 1 << 2
//	MsgAttrSent                  MsgAttr   = 1 << 3
	MsgAttrFileAttached          MsgAttr   = 1 << 4
//	MsgAttrInTransit             MsgAttr   = 1 << 5
//	MsgAttrOrphan                MsgAttr   = 1 << 6
//	MsgAttrKillSent              MsgAttr   = 1 << 7
//	MsgAttrLocal                 MsgAttr   = 1 << 8
//	MsgAttrHoldForPickup         MsgAttr   = 1 << 9
//	MsgAttr???                   MsgAttr = 1 << 10
//	MsgAttrFileRequest           MsgAttr = 1 << 11
//	MsgAttrReturnReceiptRequest  MsgAttr= 1 << 12
//	MsgAttrIsReturnReceipt       MsgAttr= 1 << 13
//	MsgAttrAuditRequest          MsgAttr= 1 << 14
//	MasgAttrFileUpdateReq        MsgAttr= 1 << 15
)