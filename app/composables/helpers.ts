import type { BackendErrorResponse } from "~/types/types";

const helpers = {
  isBackendError(err: unknown): err is { data: BackendErrorResponse } {
    return (
      typeof err === "object" &&
      err !== null &&
      "data" in err &&
      typeof (err as { data: unknown }).data === "object" &&
      (err as { data: unknown }).data !== null &&
      "message" in (err as { data: { message: string } }).data
    );
  },
  getBackendErrorMessage(err: unknown): string | null {
    if (helpers.isBackendError(err)) {
      return err.data.message;
    }
    return null;
  },
};

export default helpers;
