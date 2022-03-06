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
import { LoginFormType } from "../../types/forms";
type RegisterFormProps = {
  onSubmit: (data: any) => void;
};

const LoginForm: FC<RegisterFormProps> = ({ onSubmit }: RegisterFormProps) => {
  const [show, setShow] = useState(false);
  const handleShow = () => setShow(!show);

  const {
    handleSubmit,
    register,
    formState: { errors, isSubmitting },
  } = useForm<LoginFormType>();

  return (
    <Box w="75%" ml="auto" mr="auto">
      <form onSubmit={handleSubmit(onSubmit)}>
        <FormControl isInvalid={!!errors.username}>
          <FormLabel htmlFor="username" mt={2}>
            Brukernavn
          </FormLabel>
          <Input
            id="username"
            placeholder="Brukernavn"
            {...register("username", {
              required: "Påkrevd felt",
            })}
          />
          <FormErrorMessage>
            {errors.username && errors.username.message}
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

        <Button
          mt={4}
          colorScheme="teal"
          isLoading={isSubmitting}
          type="submit"
        >
          Logg inn
        </Button>
      </form>
    </Box>
  );
};

export default LoginForm;