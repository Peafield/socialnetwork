import React from "react";

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

  switch (type) {
    case "success":
      backgroundColor = "green";
      break;
    case "error":
      backgroundColor = "tomato";
      break;
    case "warning":
      backgroundColor = "yellow";
      break;
    default:
      backgroundColor = "lightgray";
  }

  const style = {
    position: "fixed" as "fixed",
    bottom: "10px",
    left: "10px",
    backgroundColor: backgroundColor,
    color: "white",
    padding: "10px 20px",
    borderRadius: "5px",
    boxShadow:
      "rgba(50, 50, 93, 0.25) 0px 2px 5px -1px, rgba(0, 0, 0, 0.3) 0px 1px 3px -1px",
    cursor: "pointer",
  };

  return (
    <div style={style} onClick={onClose}>
      {message}
    </div>
  );
};

export default Snackbar;
