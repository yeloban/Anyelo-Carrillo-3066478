public class variable {
    public static void main(String[] args) {
        int a = 10;
        int b = 4;

        // Multiplicación
        System.out.println("La primera variable es: " + a);
        System.out.println("El signo de la operación es: *");
        System.out.println("La segunda variable es: " + b);
        int c = a * b;
        System.out.println("El resultado es: " + c);
        System.out.println("Tipo de dato: " + ((Object) c).getClass().getSimpleName());

        // División
        System.out.println("La primera variable es: " + a);
        System.out.println("El signo de la operación es: /");
        System.out.println("La segunda variable es: " + b);
        double cDiv = (double) a / b; // Conversión a double para división precisa
        System.out.println("El resultado es: " + cDiv);
        System.out.println("Tipo de dato: " + ((Object) cDiv).getClass().getSimpleName());

        // Suma
        System.out.println("La primera variable es: " + a);
        System.out.println("El signo de la operación es: +");
        System.out.println("La segunda variable es: " + b);
        int cSuma = a + b;
        System.out.println("El resultado es: " + cSuma);
        System.out.println("Tipo de dato: " + ((Object) cSuma).getClass().getSimpleName());

        // Resta
        System.out.println("La primera variable es: " + a);
        System.out.println("El signo de la operación es: -");
        System.out.println("La segunda variable es: " + b);
        int cResta = a - b;
        System.out.println("El resultado es: " + cResta);
        System.out.println("Tipo de dato: " + ((Object) cResta).getClass().getSimpleName());

        // Potencia (usando Math.pow)
        System.out.println("La primera variable es: " + a);
        System.out.println("El signo de la operación es: **");
        System.out.println("La segunda variable es: " + b);
        double cPotencia = Math.pow(a, b); // Math.pow devuelve un double
        System.out.println("El resultado es: " + cPotencia);
        System.out.println("Tipo de dato: " + ((Object) cPotencia).getClass().getSimpleName());

        // Módulo
        System.out.println("La primera variable es: " + a);
        System.out.println("El signo de la operación es: %");
        System.out.println("La segunda variable es: " + b);
        int cModulo = a % b;
        System.out.println("El resultado es: " + cModulo);
        System.out.println("Tipo de dato: " + ((Object) cModulo).getClass().getSimpleName());
    }
}
