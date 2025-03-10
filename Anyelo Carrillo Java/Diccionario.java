import java.util.HashMap;
import java.util.Map;

public class Diccionario {
    public static void main(String[] args) {
        Map<String, String> coche = new HashMap<>();
        coche.put("marca", "Ford");
        coche.put("color", "rojo");
        coche.put("modelo", "sedan");
        coche.put("placa", "ROS227");

        System.out.println(coche.get("color"));
        coche.put("color", "verde");
        System.out.println(coche.get("color"));
        System.out.println(coche.get("marca"));
        coche.put("marca", "Renault");
        System.out.println(coche.get("marca"));
        System.out.println(coche);
    }
}
