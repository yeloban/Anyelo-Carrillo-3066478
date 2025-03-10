import java.util.ArrayList;
import java.util.List;

public class Funciones {

    // Variables globales
    static int a = 5;
    static int b = 8;

    public static void main(String[] args) {
        // Parte 1: Multiplicación de elementos de dos listas
        int[] a = {1, 2, 3, 4, 5};
        int[] b = {6, 7, 8, 9, 10};
        List<Integer> c = new ArrayList<>();

        for (int i = 0; i < a.length; i++) {
            c.add(a[i] * b[i]);
        }
        System.out.println(c);

        // Parte 2: Funciones
        mostrarTexto(); // Llamada a la función mostrarTexto
        multiplicar();  // Llamada a la función multiplicar

        // Llamada a la función multiplicar con variables globales
        multiplicarConGlobales();

        // Llamada a la función multiplicarConReturn
        int resultado = multiplicarConReturn();
        System.out.println(resultado + 5);

        // Llamada a la función validarDato
        boolean dato = validarDato(a);
        System.out.println(dato);
    }

    // Función para mostrar texto
    static void mostrarTexto() {
        System.out.println("hola");
    }

    // Función para multiplicar dos números locales
    static void multiplicar() {
        int a = 5;
        int b = 8;
        System.out.println(a * b);
    }

    // Función para multiplicar usando variables globales
    static void multiplicarConGlobales() {
        System.out.println(a * b);
    }

    // Función para multiplicar y devolver el resultado
    static int multiplicarConReturn() {
        int a = 5;
        int b = 8;
        return a * b;
    }

    // Función para validar un dato
    static boolean validarDato(int a) {
        return a == 5;
    }
}




