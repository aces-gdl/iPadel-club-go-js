import { useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export const ResponseInterceptor = () => {
  const navigate = useNavigate();
  const interceptorId = useRef(null);

  useEffect(() => {
    interceptorId.current = axios.interceptors.response.use(
      (response) => response,
      (error) => {
        console.log("Interceptor caught an error:", error);

        if (error.response) {
          console.log("Error status:", error.response.status);
          console.log("Error data:", error.response.data);

          switch (error.response.status) {
            case 401:
              console.log("Unauthorized access detected. Redirecting to login...");
              localStorage.removeItem("user");
              navigate("/auth/login");
              return Promise.reject(error);
            default:
              return Promise.reject(error);
          }
        } else if (error.request) {
          console.log("No response received:", error.request);
        } else {
          console.log("Error setting up request:", error.message);
        }

        return Promise.reject(error);
      }
    );

    // Cleanup function
    return () => {
      axios.interceptors.response.eject(interceptorId.current);
    };
  }, [navigate]);

  return null;
};
