import axios from "axios";

export async function fetchGreet() {
    const response = await axios.get("http://localhost:8000/go");

    if (response.data) {
        return response.data;
    }
}