import React, { ReactNode, useEffect, useRef } from "react";
import styles from "./Modal.module.css";

interface ModalProps {
  children: ReactNode;
  open: boolean;
  onClose: () => void;
}

const Modal: React.FC<ModalProps> = ({ children, open, onClose }) => {
  const modalRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const outsideClickListener = (event: MouseEvent) => {
      if (
        modalRef.current &&
        !modalRef.current.contains(event.target as Node)
      ) {
        onClose();
      }
    };

    if (open) {
      document.addEventListener("mousedown", outsideClickListener);
    } else {
      document.removeEventListener("mousedown", outsideClickListener);
    }

    return () => {
      document.removeEventListener("mousedown", outsideClickListener);
    };
  }, [open, onClose]);

  if (!open) return null;

  return (
    <div className={styles.darkbackground}>
      <div ref={modalRef} className={styles.modal}>
        {children}
      </div>
    </div>
  );
};

export default Modal;
