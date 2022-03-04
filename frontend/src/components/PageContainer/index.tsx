import React, { FC } from "react";
import { Box, Heading } from "@chakra-ui/react";

type PageContainerProps = {
  children: React.ReactNode;
  title: string;
};

const PageContainer: FC<PageContainerProps> = ({
  children,
  title,
}: PageContainerProps) => {
  return (
    <Box w="80%" marginLeft="auto" marginRight="auto" textAlign="center">
      <Heading as="h1" mb={4} mt={4}>
        {title}
      </Heading>
      {children}
    </Box>
  );
};

export default PageContainer;
