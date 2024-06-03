package wrapper

import (
	"github.com/mariemalysheva/tokenized-reinvestment/internal/models"
)

const blockScannerUrl = "https://testnet.bscscan.com/tx/"

func RepackAddUserReq(req []UserReq) (users []models.Client) {
	for _, u := range req {
		users = append(users, models.Client{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Savings:   u.Savings,
		})
	}
	return users
}

func RepackReinvestmentRecord(rec models.ReinvestmentRecord) ReinvestmentRecord {
	return ReinvestmentRecord{
		Reinvestment: Reinvestment{
			ID:    rec.Reinvestment.ID,
			Asset: rec.Reinvestment.Asset.Hex(),
			Rate:  rec.Reinvestment.Rate,
			Price: rec.Reinvestment.Price,
			Start: rec.Reinvestment.Start,
			End:   rec.Reinvestment.End,
		},
		Savings: rec.Savings,
		Amount:  rec.Amount,
	}
}

func RepackReinvestment(rec models.Reinvestment) Reinvestment {
	return Reinvestment{
		ID:    rec.ID,
		Asset: rec.Asset.Hex(),
		Rate:  rec.Rate,
		Price: rec.Price,
		Start: rec.Start,
		End:   rec.End,
	}
}

func RepackRecords(recs []models.ReinvestmentRecord) []ReinvestmentRecord {
	resp := make([]ReinvestmentRecord, 0, len(recs))
	for _, rec := range recs {
		resp = append(resp, RepackReinvestmentRecord(rec))
	}
	return resp
}

func RepackReinvestmentPeriods(recs []models.Reinvestment) []Reinvestment {
	resp := make([]Reinvestment, 0, len(recs))
	for _, rec := range recs {
		resp = append(resp, RepackReinvestment(rec))
	}
	return resp
}

func RepackTxHashResp(txHash string) TxResp {
	return TxResp{TxHash: blockScannerUrl + txHash}
}
