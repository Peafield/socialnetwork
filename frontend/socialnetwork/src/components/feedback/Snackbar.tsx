import React, { CSSProperties } from "react";

import {
  FiAlertTriangle,
  FiCheck,
  FiAlertOctagon,
  FiAlertCircle,
} from "react-icons/fi";

import { AiOutlineClose } from "react-icons/ai";

type SnackbarType = "success" | "error" | "warning";

interface SnackbarProps {
  open: boolean;
  onClose: () => void;
  message: string;
  type: SnackbarType;
}

const Snackbar: React.FC<SnackbarProps> = ({
  open,
  onClose,
  message,
  type,
}) => {
  if (!open) return null;

  let backgroundColor: string;
  let Icon: React.ComponentType;

  switch (type) {
    case "success":
      backgroundColor = "green";
      Icon = FiCheck;
      break;
    case "error":
      backgroundColor = "tomato";
      Icon = FiAlertTriangle;
      break;
    case "warning":
      backgroundColor = "#ffbf47";
      Icon = FiAlertOctagon;
      break;
    default:
      backgroundColor = "lightgray";
      Icon = FiAlertCircle;
  }

  const style = {
    position: "fixed" as "fixed",
    top: "20px",
    left: "50%",
    transform: "translateX(-50%)",
    backgroundColor: backgroundColor,
    color: "white",
    padding: "10px 20px",
    borderRadius: "5px",
    boxShadow:
      "rgba(50, 50, 93, 0.25) 0px 2px 5px -1px, rgba(0, 0, 0, 0.3) 0px 1px 3px -1px",
    cursor: "pointer",
    transition: "all 0.4s ease",
  };

  const iconStyle: CSSProperties = {
    marginRight: "8px",
    verticalAlign: "middle",
  };
  const closeStyle: CSSProperties = {
    marginLeft: "8px",
    verticalAlign: "middle",
  };

  return (
    <div style={style} onClick={onClose}>
      <span style={iconStyle}>
        <Icon />
      </span>
      {message}
      <span style={closeStyle}>
        <AiOutlineClose />
      </span>
    </div>
  );
};

export default Snackbar;
