import express from 'express';

import noteRtr from './note';

const router = express.Router();

router.use('/note', noteRtr);

export default router;
