package ngapConvert

import (
	"fmt"

	"github.com/omec-project/ngap/ngapType"
	"github.com/omec-project/openapi/models"
)

func NrCgiToModels(nrcgi ngapType.NRCGI) models.Ncgi {
	var modelsNrCgi models.Ncgi

	plmnID := PlmnIdToModels(nrcgi.PLMNIdentity)
	modelsNrCgi.PlmnId = &plmnID
	//modelsNrCgi.NrCellId = hex.EncodeToString(nrcgi.NRCellIdentity.Value)
	modelsNrCgi.NrCellId = BitStringToHex(&nrcgi.NRCellIdentity.Value)
	return modelsNrCgi
}

func NrCgiToNgap(ncgi models.Ncgi) ngapType.NRCGI {
	var ngapNrCgi ngapType.NRCGI
	fmt.Println("ncgi.NrCellId:", ncgi.NrCellId, "EOS")
	fmt.Println("len ncgi.NrCellId:", len(ncgi.NrCellId))

	ngapNrCgi.PLMNIdentity = PlmnIdToNgap(*ncgi.PlmnId)
	ngapNrCgi.NRCellIdentity.Value = HexToBitString(ncgi.NrCellId, 36)
	return ngapNrCgi
}
