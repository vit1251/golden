package msg

import (
	"log"
	"testing"
)

func TestMessageLineParser1(t *testing.T) {
	mlp := NewMessageLineParser()
	ml := mlp.Parse(" VS> Добрый день!")

	log.Printf("ml = %+v", ml)

	if ml.QuoteStart != " " {
		t.Errorf("Wrong QuoteStart value %+v", ml.QuoteStart)
	}
	if ml.QuoteAuthor != "VS" {
		t.Errorf("Wrong QuoteAuthor value %+v", ml.QuoteAuthor)
	}
	if ml.QuoteLevel != 1 {
		t.Errorf("Wrong QuteLevel value %+v", ml.QuoteLevel)
	}
	if ml.QuoteLine != " Добрый день!" {
		t.Errorf("Wrong QuteLine value %+v", ml.QuoteLine)
	}

}

func TestMessageLineParser2(t *testing.T) {

	mlp := NewMessageLineParser()
	ml := mlp.Parse("Добрый день!")

	log.Printf("ml = %+v", ml)

	if ml.QuoteAuthor != "" {
		t.Errorf("Wrong QuoteAuthor value %+v", ml.QuoteAuthor)
	}
	if ml.QuoteLevel != 0 {
		t.Errorf("Wrong QuteLevel value %+v", ml.QuoteLevel)
	}
	if ml.QuoteLine != "Добрый день!" {
		t.Errorf("Wrong QuteLine value %+v", ml.QuoteLine)
	}

}

func TestMessageLineParser3(t *testing.T) {
	mlp := NewMessageLineParser()
	ml := mlp.Parse(" 25 августа 1995 написал Vitold Sedyshev -> Alexander Kirilov")

	log.Printf("ml = %+v", ml)

	if ml.QuoteAuthor != "" {
		t.Errorf("Wrong QuoteAuthor value %+v", ml.QuoteAuthor)
	}
	if ml.QuoteLevel != 0 {
		t.Errorf("Wrong QuteLevel value %+v", ml.QuoteLevel)
	}
	if ml.QuoteLine != " 25 августа 1995 написал Vitold Sedyshev -> Alexander Kirilov" {
		t.Errorf("Wrong QuteLine value %+v", ml.QuoteLine)
	}

}
