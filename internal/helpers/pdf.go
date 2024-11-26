package helpers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Jollynjose/sistema-viatico-backend/internal/infrastructure/db"
	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GeneratePDF(travelExpense *db.TravelExpense, fuel *db.Fuel) ([]byte, error) {
	m, err := getMaroto(travelExpense, fuel)

	if err != nil {
		return nil, err
	}

	document, err := m.Generate()
	if err != nil {
		return nil, err
	}

	return document.GetBytes(), nil
}

func getMaroto(travelExpense *db.TravelExpense, fuel *db.Fuel) (core.Maroto, error) {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(15).
		WithRightMargin(10).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader(travelExpense))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRow(7,
		text.NewCol(12, "Calculo de Consumo", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
		}),
	)

	m.AddRows(getTransactions(travelExpense)...)

	m.AddRow(7,
		text.NewCol(12, "Calculo de Combustible", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
		}),
	)

	combustiblesRow, err := getCombustibles(travelExpense, fuel)

	if err != nil {
		return nil, err
	}

	m.AddRows(combustiblesRow...)

	m.AddRow(6)

	m.AddRows(getPeaje(travelExpense)...)

	m.AddRow(6)

	m.AddRows(getMontoTotal(travelExpense)...)

	m.AddRow(6)

	m.AddRow(12,
		col.New(1),
		signature.NewCol(5, "Solicitado Por:"),
		signature.NewCol(5, "Aprobado Por: "),
		col.New(1),
	)

	m.AddRow(12,
		col.New(1),
		signature.NewCol(5, "Revisado Por: "),
		signature.NewCol(5, "Autorizado Por: "),
		col.New(1),
	)

	m.AddRow(12,
		col.New(3),
		signature.NewCol(6, "Recibido conforme: "),
		col.New(3),
	)

	return m, nil
}

func getTransactions(travelExpense *db.TravelExpense) []core.Row {
	rows := []core.Row{
		row.New(4).Add(
			text.NewCol(1, "Fecha", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Nombre", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Puesto", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Hospedaje", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Desayuno", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Almuerzo", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Cena", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Pasaje", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Total", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		).WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 0, Green: 200, Blue: 0}}),
	}

	var contentsRow []core.Row

	for i, content := range travelExpense.UserTravelHistory {
		accommadation := content.JobPositionHistory.Accommodation
		breakfast := content.JobPositionHistory.BreakFast
		lunch := content.JobPositionHistory.Lunch
		dinner := content.JobPositionHistory.Dinner

		if !content.IsAccommodationApplied {
			accommadation = 0
		}

		if !content.IsBreakfastApplied {
			breakfast = 0
		}

		if !content.IsLunchApplied {
			lunch = 0
		}

		if !content.IsDinnerApplied {
			dinner = 0
		}

		total := accommadation + breakfast + lunch + dinner + content.PassagePrice

		r := row.New(4).Add(
			text.NewCol(1, content.CreatedAt.Format("2006-01-02"), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, fmt.Sprintln(content.User.FirstName, content.User.LastName), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, content.User.JobPosition.Name, props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, strconv.FormatFloat(accommadation, 'f', -1, 64), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, strconv.FormatFloat(breakfast, 'f', -1, 64), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, strconv.FormatFloat(lunch, 'f', -1, 64), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, strconv.FormatFloat(dinner, 'f', -1, 64), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, strconv.FormatFloat(content.PassagePrice, 'f', -1, 64), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, strconv.FormatFloat(total, 'f', -1, 64), props.Text{Size: 8, Align: align.Center}),
		)
		if i%2 == 1 {
			gray := getGrayColor()
			r.WithStyle(&props.Cell{BackgroundColor: gray})
		}

		contentsRow = append(contentsRow, r)
	}

	rows = append(rows, contentsRow...)

	return rows
}

func getCombustibles(travelExpense *db.TravelExpense, fuel *db.Fuel) ([]core.Row, error) {
	rows := []core.Row{
		row.New(4).Add(
			text.NewCol(2, "Combustible", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Kilometraje", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Galones Consumidos", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Precio por galon RD$", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Total RD$", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		).WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 0, Green: 200, Blue: 0}}),
	}

	var contentsRow []core.Row

	galonesConsumidos := travelExpense.Route.TotalKms / 30

	galonesPrecio := float64(travelExpense.Route.TotalKms) * travelExpense.FuelHistory.Price

	r := row.New(4).Add(
		text.NewCol(2, fuel.Name, props.Text{Size: 8, Align: align.Center}),
		text.NewCol(2, IntToString(travelExpense.Route.TotalKms), props.Text{Size: 8, Align: align.Center}),
		text.NewCol(3, IntToString(galonesConsumidos), props.Text{Size: 8, Align: align.Center}),
		text.NewCol(3, FloatToString(travelExpense.FuelHistory.Price), props.Text{Size: 8, Align: align.Center}),
		text.NewCol(2, FloatToString(galonesPrecio), props.Text{Size: 8, Align: align.Center}),
	)

	contentsRow = append(contentsRow, r)

	rows = append(rows, contentsRow...)

	return rows, nil
}

func getPeaje(travelExpense *db.TravelExpense) []core.Row {
	rows := []core.Row{
		row.New(4).Add(
			col.New(6),
			text.NewCol(2, "C/peaje", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).
				WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 0, Green: 200, Blue: 0}}),
			text.NewCol(2, "Precio Peaje", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).
				WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 0, Green: 200, Blue: 0}}),
			text.NewCol(2, "Total RD$", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).
				WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 0, Green: 200, Blue: 0}}),
		),
	}

	var contentsRow []core.Row
	contents := travelExpense.Toll
	/*for i := 0; i < 8; i++ {
	    contents = append(contents, contents...)
	}*/

	for _, content := range contents {
		r := row.New(4).Add(
			col.New(6),
			text.NewCol(2, IntToString(content.Order), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, FloatToString(content.Price), props.Text{Size: 8, Align: align.Center}),
			text.NewCol(2, FloatToString(content.Price), props.Text{Size: 8, Align: align.Center}),
		)

		contentsRow = append(contentsRow, r)
	}

	rows = append(rows, contentsRow...)

	return rows
}

func getPageHeader(travelExpense *db.TravelExpense) core.Row {
	return row.New(25).Add(
		col.New(6).Add(
			text.New(fmt.Sprintf("Fecha de solicitud: %s", FormatTimeToYYYYMMDD(travelExpense.SolicitudeDate)), props.Text{
				Top: 0,
			}),
			text.New(fmt.Sprintf("Dependencia: %s", travelExpense.Dependency), props.Text{
				Top: 5,
			}),

			text.New(fmt.Sprintf("Transporte: %s", travelExpense.TransporteType), props.Text{
				Top: 10,
			}),
			text.New(fmt.Sprintf("Hora de Salida: %s", FormatTimeToHHMMAMOrPM(travelExpense.DepartureDate)), props.Text{
				Top: 15,
			}),
		),
		col.New(6).Add(
			text.New(fmt.Sprintf("Motivo de salida: %s", travelExpense.VisitMotivation), props.Text{
				Top: 0,
			}),
			text.New(fmt.Sprintf("Lugar de Partida: %s", travelExpense.Route.StartingPointProvince.Name), props.Text{
				Top: 5,
			}),

			text.New(fmt.Sprintf("Lugar de visita: %s", travelExpense.Route.FinalDestinationProvince.Name), props.Text{
				Top: 10,
			}),
			text.New(fmt.Sprintf("Hora de llegada: %s", FormatTimeToHHMMAMOrPM(travelExpense.ArrivalDate)), props.Text{
				Top: 15,
			}),
		),
	)
}

func getMontoTotal(travel *db.TravelExpense) []core.Row {
	rows := []core.Row{
		row.New(4).Add(
			col.New(8),
			text.NewCol(4, "Monto Total RD$", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}).
				WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 0, Green: 200, Blue: 0}}),
		),
	}

	var contentsRow []core.Row

	r := row.New(4).Add(
		col.New(8),
		text.NewCol(4, FloatToString(travel.TotalPrice), props.Text{Size: 8, Align: align.Center}),
	)

	contentsRow = append(contentsRow, r)

	rows = append(rows, contentsRow...)

	return rows
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}
