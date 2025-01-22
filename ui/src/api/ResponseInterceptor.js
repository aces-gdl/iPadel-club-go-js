import { useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import api from "./api";
import axios from "axios";

export const ResponseInterceptor = () => {
  const navigate = useNavigate();

  const interceptorId = useRef(null);

  useEffect(() => {
    interceptorId.current = axios.interceptors.response.use(
      undefined,
      (error) => {
        switch (error.response.status) {
          case 401:
            localStorage.removeItem("user");
            navigate("/");
            break;
          default:
            return Promise.reject(error);
        }
      }
    );
  }, [navigate]);

  return null;
};
