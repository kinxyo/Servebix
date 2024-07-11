import axios from "axios";

export async function fetchGreet() {
    const response = await axios.get("http://192.168.1.6:8000/go");

    if (response.data) {
        return response.data;
    }
}

export async function ShowAllPatients(){
    const response = await axios.get("http://192.168.1.6:8000/patient/all");

        if (response.data) {
            return response.data;
        }
}
