import SimpleCenterLayout from "@/libs/globals/components/layouts/SimpleCenterLayout";
import SignUpForm from "@/libs/routes/auth/SignUpForm";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/auth/sign-up")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <SimpleCenterLayout>
      <SignUpForm />
    </SimpleCenterLayout>
  );
}
