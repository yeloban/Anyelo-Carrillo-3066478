import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class listas {
    public static void main(String[] args) {
        // Crear una lista de días de la semana
        List<String> lista = new ArrayList<>(Arrays.asList("Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado"));

        // Imprimir el elemento en la posición 4 (índice 4)
        System.out.println(lista.get(4)); // Imprime "viernes"

        // Imprimir la lista completa
        System.out.println(lista); // Imprime toda la lista

        // Imprimir una sublista desde el índice 0 hasta el 3 (exclusivo)
        System.out.println(lista.subList(0, 3)); // Imprime ["Lunes", "Martes", "Miercoles"]

        // Crear una lista heterogénea (mezcla de tipos)
        List<Object> listaHeterogenea = new ArrayList<>(Arrays.asList(
            "Lunes", "Martes", "Miercoles", "Jueves", "viernes", "sabado", 1, 2, 3,
            Arrays.asList("Esteban", 0.2, 2.25, true)
        ));

        // Imprimir una sublista desde el índice 0 hasta el 4 (exclusivo)
        System.out.println(listaHeterogenea.subList(0, 4)); // Imprime ["Lunes", "Martes", "Miercoles", "Jueves"]

        // Imprimir una sublista de la lista anidada en la posición 9
        List<Object> subLista = (List<Object>) listaHeterogenea.get(9); // Obtener la lista anidada
        System.out.println(subLista.subList(0, 3)); // Imprime ["Esteban", 0.2, 2.25]
    }
}
