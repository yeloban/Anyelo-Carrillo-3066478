import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        // Lista de nombres
        String[] nombres = {"Esteban", "Hans", "Jhon", "Juan Pablo \n\n"};
        for (String nombre : nombres) {
            System.out.println(nombre);
        }

        // Lista de personas (diccionarios)
        List<Map<String, Object>> Personas = new ArrayList<>();

        Map<String, Object> a = new HashMap<>();
        a.put("nombre", "Esteban");
        a.put("Edad", 28);

        Map<String, Object> b = new HashMap<>();
        b.put("nombre", "Hans");
        b.put("Edad", 27);

        Map<String, Object> c = new HashMap<>();
        c.put("nombre", "Jhon");
        c.put("Edad", 41);

        Map<String, Object> d = new HashMap<>();
        d.put("nombre", "Juan Pablo");
        d.put("Edad", 23);

        Personas.add(a);
        Personas.add(b);
        Personas.add(c);
        Personas.add(d);

        for (Map<String, Object> persona : Personas) {
            System.out.println(persona.get("nombre") + " " + persona.get("Edad"));
        }
    }
}
