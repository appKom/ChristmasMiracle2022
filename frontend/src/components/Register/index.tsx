import React, { FC, useState } from "react";
import { useForm } from "react-hook-form";
import {
  FormErrorMessage,
  FormLabel,
  FormControl,
  Input,
  Button,
  InputGroup,
  InputRightElement,
  Box,
} from "@chakra-ui/react";
import { RegisterFormType } from "../../types/forms";
type RegisterFormProps = {
  onSubmit: (data: any) => void;
};

const RegisterForm: FC<RegisterFormProps> = ({
  onSubmit,
}: RegisterFormProps) => {
  const [show, setShow] = useState(false);
  const [showVerify, setShowVerify] = useState(false);
  const handleShow = () => setShow(!show);
  const handleShowVerify = () => setShowVerify(!showVerify);

  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm<RegisterFormType>();

  const checkUsernameAvailability = (username: string): boolean => {
    return true;
  };

  const validatePassword = (value: string): boolean => {
    if (value.length < 8) {
      return false;
    }
    return true;
  };

  return (
    <Box w="75%" ml="auto" mr="auto">
      <form onSubmit={handleSubmit(onSubmit)}>
        <FormControl isInvalid={!!errors.username}>
          <FormLabel htmlFor="username" mt={2}>
            Brukernavn (dette vil vises til alle)
          </FormLabel>
          <Input
            id="username"
            placeholder="Brukernavn"
            {...register("username", {
              required: "Påkrevd felt",
              minLength: {
                value: 3,
                message: "Brukernavn må være lenger enn 3 tegn",
              },
              validate: (username: string) =>
                checkUsernameAvailability(username) ||
                "Brukernavnet er ikke ledig",
            })}
          />
          <FormErrorMessage>
            {errors.username && errors.username.message}
          </FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={!!errors.email}>
          <FormLabel htmlFor="email" mt={2}>
            Epost (brukes til kommunikasjon)
          </FormLabel>
          <Input
            id="email"
            placeholder="Epost"
            {...register("email", {
              required: "Påkrevd felt",
              pattern: {
                // eslint-disable-next-line no-useless-escape
                value: /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/,
                message: "Eposten er ikke gyldig",
              },
            })}
          />
          <FormErrorMessage>
            {errors.email && errors.email.message}
          </FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={!!errors.password}>
          <FormLabel htmlFor="password" mt={2}>
            Passord
          </FormLabel>

          <InputGroup size="md">
            <Input
              id="password"
              type={show ? "text" : "password"}
              placeholder="Passord"
              {...register("password", {
                required: "Påkrevd felt",
                validate: (password: string) =>
                  validatePassword(password) || "Passordet er ikke gyldig",
              })}
            />
            <InputRightElement width="4.5rem">
              <Button h="1.75rem" size="sm" color="black" onClick={handleShow}>
                {show ? "Hide" : "Show"}
              </Button>
            </InputRightElement>
          </InputGroup>

          <FormErrorMessage>
            {errors.password && errors.password.message}
          </FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={!!errors.verifyPassword}>
          <FormLabel htmlFor="verifyPassword" mt={2}>
            Verifiser passord
          </FormLabel>

          <InputGroup size="md">
            <Input
              id="verifyPassword"
              type={showVerify ? "text" : "password"}
              placeholder="Passord"
              {...register("verifyPassword", {
                required: "Påkrevd felt",
              })}
            />
            <InputRightElement width="4.5rem">
              <Button
                h="1.75rem"
                size="sm"
                color="black"
                onClick={handleShowVerify}
              >
                {showVerify ? "Hide" : "Show"}
              </Button>
            </InputRightElement>
          </InputGroup>

          <FormErrorMessage>
            {errors.verifyPassword && errors.verifyPassword.message}
          </FormErrorMessage>
        </FormControl>
        <Button
          mt={4}
          colorScheme="teal"
          isLoading={isSubmitting}
          type="submit"
        >
          Send inn
        </Button>
      </form>
    </Box>
  );
};

export default RegisterForm;
