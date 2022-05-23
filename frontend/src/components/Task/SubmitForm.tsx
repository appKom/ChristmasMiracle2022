import React, { FC, useState, useEffect } from "react";
import {
  Input,
  Button,
  Text,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
} from "@chakra-ui/react";
import { submitFlag } from "../../api";
import { useRecoilState } from "recoil";
import { tokenState } from "../../state/auth";

type SubmitFormProps = {
  id: number;
  modalRef: React.RefObject<HTMLButtonElement>;
};

const SubmitForm: FC<SubmitFormProps> = ({ id, modalRef }: SubmitFormProps) => {
  const [flag, setFlag] = useState("");
  const [response, setResponse] = useState("");
  const { isOpen, onOpen, onClose } = useDisclosure();

  const [token] = useRecoilState(tokenState);

  useEffect(() => {
    if (modalRef.current) {
      modalRef.current.onclick = () => onOpen();
    }
  }, [modalRef, onOpen]);

  const handleSubmit = () => {
    submitFlag(token, id, flag).then((resp) => setResponse(resp));
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFlag(e.target.value);
  };

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader color="black">Send inn oppgave</ModalHeader>
        <ModalCloseButton color="black" />
        <ModalBody>
          <Input placeholder="Flagg" onChange={handleChange} color="black" />
          <Text color="black">{response}</Text>
        </ModalBody>
        <ModalFooter>
          <Button colorScheme="red" mr={3} onClick={onClose}>
            Avbryt
          </Button>
          <Button color="teal" ml="5px" onClick={() => handleSubmit()}>
            Send inn
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
export default SubmitForm;
