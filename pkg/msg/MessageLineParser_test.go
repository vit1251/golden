package msg

import (
    "testing"
)

func TestMessageLineParser1(t *testing.T) {

	mlp := NewMessageLineParser()
	ml := mlp.Parse(" VS> Добрый день!")

        /* Ожидаемый результат */

	if ml.Author != "VS" {
		t.Errorf("Wrong QuoteAuthor value %+v", ml.Author)
	}

	if ml.QuoteLevel != 1 {
		t.Errorf("Wrong QuteLevel value %+v", ml.QuoteLevel)
	}

	if ml.Text != " Добрый день!" {
		t.Errorf("Wrong QuteLine value %+v", ml.Text)
	}

}

func TestMessageLineParser2(t *testing.T) {

	mlp := NewMessageLineParser()
	ml := mlp.Parse("Добрый день!")

        /* Ожидаемый результат */

	if ml.Author != "" {
		t.Errorf("Wrong QuoteAuthor value %+v", ml.Author)
	}

	if ml.QuoteLevel != 0 {
		t.Errorf("Wrong QuteLevel value %+v", ml.QuoteLevel)
	}

	if ml.Text != "Добрый день!" {
		t.Errorf("Wrong QuteLine value %+v", ml.Text)
	}

}

func TestMessageLineParser3(t *testing.T) {

	mlp := NewMessageLineParser()
	ml := mlp.Parse(" 25 августа 1995 написал Vitold Sedyshev -> Alexander Kirilov")

        /* Ожидаемый результат */

	if ml.Author != "" {
		t.Errorf("Wrong QuoteAuthor value %+v", ml.Author)
	}

	if ml.QuoteLevel != 0 {
		t.Errorf("Wrong QuteLevel value %+v", ml.QuoteLevel)
	}

	if ml.Text != " 25 августа 1995 написал Vitold Sedyshev -> Alexander Kirilov" {
		t.Errorf("Wrong QuteLine value %+v", ml.Text)
	}

}
