import React, { ReactNode, useEffect, useRef } from "react";
import styles from "./SideModal.module.css";

interface SideModalProps {
  children: ReactNode;
  open: boolean;
  onClose: () => void;
}

const SideModal: React.FC<SideModalProps> = ({ children, open, onClose }) => {
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
    <div ref={modalRef} className={styles.sidemodal}>
      {children}
    </div>
  );
};

export default SideModal;
