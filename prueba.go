
package paquetes

import (
    "testing"
)

func TestSuma(t *testing.T) {
    resultado := Suma(2, 3)
    esperado := 5
    if resultado != esperado {
        t.Errorf("Suma(2, 3) = %d; se esperaba %d", resultado, esperado)
    }
}
