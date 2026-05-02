import { useState, useEffect, useCallback } from 'react';

/**
 * Custom hook for making API requests with loading/error/data states.
 *
 * @param {string} url      — The endpoint URL (relative or absolute).
 * @param {object} options  — Optional config: { method, body, immediate, headers }.
 * @returns {{ data, error, loading, execute }}
 *
 * Usage:
 *   const { data, loading, error } = useApi('/api/art-of-the-day');
 *   const { data, execute }        = useApi('/api/generate-music', { immediate: false });
 */
export function useApi(url, options = {}) {
  const {
    method    = 'GET',
    body      = null,
    immediate = true,
    headers   = { 'Content-Type': 'application/json' },
  } = options;

  const [data, setData]       = useState(null);
  const [error, setError]     = useState(null);
  const [loading, setLoading] = useState(false);

  /**
   * Execute the API call. Accepts an optional override body
   * so callers can pass dynamic payloads (e.g., genre selection).
   */
  const execute = useCallback(async (overrideBody = null) => {
    setLoading(true);
    setError(null);
    setData(null);

    try {
      const fetchOptions = {
        method,
        headers,
      };

      const payload = overrideBody ?? body;
      if (payload) {
        fetchOptions.body = JSON.stringify(payload);
      }

      const response = await fetch(url, fetchOptions);

      if (!response.ok) {
        throw new Error(`API error: ${response.status} ${response.statusText}`);
      }

      const json = await response.json();
      setData(json);
      return json;
    } catch (err) {
      setError(err.message);
      return null;
    } finally {
      setLoading(false);
    }
  }, [url, method, body, headers]);

  /* Auto-fire on mount when immediate is true */
  useEffect(() => {
    if (immediate) {
      execute();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return { data, error, loading, execute };
}

export default useApi;
