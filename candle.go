package techan

import (
	"fmt"
	"strings"

	"github.com/sdcoffey/big"
)

// Candle represents basic market information for a security over a given time period
type Candle struct {
	Period     TimePeriod
	OpenPrice  big.Decimal `json:",string"`
	ClosePrice big.Decimal `json:",string"`
	MaxPrice   big.Decimal `json:",string"`
	MinPrice   big.Decimal `json:",string"`
	Volume     big.Decimal `json:",string"`
	BuyVolume  big.Decimal `json:",string"`
	SellVolume big.Decimal `json:",string"`
	TradeCount uint
}

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period:     period,
		OpenPrice:  big.ZERO,
		ClosePrice: big.ZERO,
		MaxPrice:   big.ZERO,
		MinPrice:   big.ZERO,
		Volume:     big.ZERO,
		BuyVolume:  big.ZERO,
		SellVolume: big.ZERO,
	}
}

// AddTrade adds a trade to this candle. It will determine if the current price is higher or lower than the min or max
// price and increment the tradecount.
func (c *Candle) AddTrade(tradeAmount, tradePrice big.Decimal, side OrderSide) {
	if c.OpenPrice.Zero() {
		c.OpenPrice = tradePrice
	}
	c.ClosePrice = tradePrice

	if c.MaxPrice.Zero() {
		c.MaxPrice = tradePrice
	} else if tradePrice.GT(c.MaxPrice) {
		c.MaxPrice = tradePrice
	}

	if c.MinPrice.Zero() {
		c.MinPrice = tradePrice
	} else if tradePrice.LT(c.MinPrice) {
		c.MinPrice = tradePrice
	}

	if c.Volume.Zero() {
		c.Volume = tradeAmount
	} else {
		c.Volume = c.Volume.Add(tradeAmount)
	}

	switch {
	case side == BUY:
		if c.BuyVolume.Zero() {
			c.BuyVolume = tradeAmount
		} else {
			c.BuyVolume = c.BuyVolume.Add(tradeAmount)
		}
	case side == SELL:
		if c.SellVolume.Zero() {
			c.SellVolume = tradeAmount
		} else {
			c.SellVolume = c.SellVolume.Add(tradeAmount)
		}
	}

	c.TradeCount++
}

// Strong assumption that the "n" candle is new and AFTER the start of the first
func (c *Candle) AddCandle(n *Candle) {
	if n.Period.Start.Before(c.Period.Start) {
		return // do nothing
	}

	if n.MaxPrice.GT(c.MaxPrice) {
		c.MaxPrice = n.MaxPrice.Frac(1.0)
	}

	if n.MinPrice.LT(c.MinPrice) {
		c.MinPrice = n.MinPrice.Frac(1.0)
	}

	c.Volume.Add(n.Volume)
	c.BuyVolume.Add(n.BuyVolume)
	c.SellVolume.Add(n.SellVolume)

	c.ClosePrice = n.ClosePrice.Frac(1.0)
	c.Period.End = n.Period.End

	c.TradeCount += n.TradeCount
}

func (c *Candle) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`
Time:	%s
Open:	%s
Close:	%s
High:	%s
Low:	%s
Volume:	%s
	`,
		c.Period,
		c.OpenPrice.FormattedString(7),
		c.ClosePrice.FormattedString(7),
		c.MaxPrice.FormattedString(7),
		c.MinPrice.FormattedString(7),
		c.Volume.FormattedString(7),
	))
}
